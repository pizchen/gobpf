package crt

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/pizchen/gobpf/pkg/tracer"
)

const (
	TimeFormat = "2006-01-02 15:04:05.000"
)

type LogLevel uint

const (
	Error LogLevel = iota
	Major
	Minor
	Info
	Debug
	Always
)

var (
	logLevelPrefix = map[LogLevel]string{
		Error:  "[ERROR] ",
		Major:  "[MAJOR] ",
		Minor:  "[MINOR] ",
		Info:   "[INFOR] ",
		Debug:  "[DEBUG] ",
		Always: "[*****] ",
	}

	Args struct {
		Cycles   uint
		Interval uint
		LifeTime uint
		Debug    uint
		Justime  uint
	}
)

func debug(reqLevel LogLevel, msg interface{}) {
	if reqLevel <= LogLevel(Args.Debug) || reqLevel == Always {
		log.Println(fmt.Sprintf("%v%v", logLevelPrefix[reqLevel], msg))
	}
}

type AppTracer interface {
	FlagParse()
	Begin() (*tracer.Tracer, error)
	DoCycle(t *tracer.Tracer, start time.Time)
	End(t *tracer.Tracer)
}

func Run(app AppTracer) {

	flag.UintVar(&Args.LifeTime, "l", 3600*10, "seconds of program life time")
	flag.UintVar(&Args.Interval, "i", 1000, "milli-seconds of report interval")
	flag.UintVar(&Args.Cycles, "n", 0, "number of intervals to exit")
	flag.UintVar(&Args.Justime, "j", 0, "number of micro seconds to justify")
	flag.UintVar(&Args.Debug, "d", uint(Error), "debug switch (0[E],1[Ma],2[Mi],3[I],4[D],5[A])")

	app.FlagParse()

	debug(Debug, "crt.Run enters")

	// expand interval to micro seconds
	Args.Interval *= 1000

	tracer, err := app.Begin()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(
		context.Background(), time.Second*time.Duration(Args.LifeTime))
	wg := &sync.WaitGroup{}

	go func() {
		wg.Add(1)
		defer wg.Done()
		select {
		case s := <-sigChan:
			debug(Debug, fmt.Sprintf("Signal received: %v", s))
			cancel()
		case <-ctx.Done():
		}
	}()

	go func() {

		wg.Add(1)
		defer wg.Done()

		cycle := uint(0)
		remain := Args.Interval

		for {
			select {

			case <-ctx.Done():
				debug(Debug, "Main loop ctx.Done() received")
				return

			case <-time.After(time.Microsecond * time.Duration(remain)):

				start := time.Now()

				debug(Debug, "DoCycle called")
				app.DoCycle(tracer, start)

				if Args.Cycles > 0 {
					if cycle++; cycle >= Args.Cycles {
						cancel()
						return
					}
				}

				end := time.Now()
				elapsed := end.UnixNano() - start.UnixNano()
				remain = Args.Interval - uint(elapsed/1000) - Args.Justime
				if remain <= 0 {
					remain = Args.Interval
				}
			}
		}
	}()

	tracer.Start()

	<-ctx.Done()

	// Run application end jobs before tracer stops
	app.End(tracer)
	debug(Debug, "Application ended")

	tracer.Stop()

	wg.Wait()
}
