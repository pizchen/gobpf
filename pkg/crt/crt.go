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
	Initialize() (*tracer.Tracer, error)
	DoJob(t *tracer.Tracer, start time.Time)
}

func Run(app AppTracer, args *map[string]uint) {

	flag.UintVar(&Args.LifeTime, "l", 3600*10, "seconds of program life time")
	flag.UintVar(&Args.Interval, "i", 1000, "milli-seconds of report interval")
	flag.UintVar(&Args.Cycles, "n", 0, "number of intervals to exit")
	flag.UintVar(&Args.Justime, "j", 0, "number of micro seconds to justify")
	flag.UintVar(&Args.Debug, "d", uint(Error), "debug switch")

	app.FlagParse()

	if args != nil {
		for k, v := range *args {
			switch k {
			case "n":
				Args.Cycles = v
			case "i":
				Args.Interval = v
			case "l":
				Args.LifeTime = v
			case "d":
				Args.Debug = v
			case "j":
				Args.Justime = v
			}
		}
	}

	// expand interval to micro seconds
	Args.Interval *= 1000

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(
		context.Background(), time.Second*time.Duration(Args.LifeTime))
	wg := &sync.WaitGroup{}

	go func() {
		wg.Add(1)
		defer wg.Done()
		select {
		case <-sigChan:
			cancel()
		case <-ctx.Done():
		}
	}()

	tracer, err := app.Initialize()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	go func() {

		wg.Add(1)
		defer wg.Done()

		cycle := uint(0)
		remain := Args.Interval

		for {
			select {

			case <-ctx.Done():
				return

			case <-time.After(time.Microsecond * time.Duration(remain)):

				start := time.Now()

				app.DoJob(tracer, start)

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
	tracer.Stop()

	wg.Wait()
}
