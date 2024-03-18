package ksym

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const (
	KALLSYMS = "/proc/kallsyms"
	KFLTFUNC = "/sys/kernel/debug/tracing/available_filter_functions"
)

type ksymCache struct {
	sync.RWMutex
	funcs map[string]struct{}
	ksym  map[uint64]string
	addr  []uint64
}

var cache ksymCache

func init() {

	cache.funcs = make(map[string]struct{})
	cache.ksym = make(map[uint64]string)
	cache.addr = make([]uint64, 0)

	fd1, err := os.Open(KFLTFUNC)
	if err != nil {
		fmt.Println("Error opening " + KALLSYMS)
		os.Exit(1)
	}
	defer fd1.Close()

	s1 := bufio.NewScanner(fd1)
	for s1.Scan() {
		l := s1.Text()
		ar := strings.Fields(l)
		cache.funcs[ar[0]] = struct{}{}
	}

	fd2, err := os.Open(KALLSYMS)
	if err != nil {
		fmt.Println("Error opening " + KALLSYMS)
		os.Exit(1)
	}
	defer fd2.Close()

	s2 := bufio.NewScanner(fd2)
	for s2.Scan() {
		l := s2.Text()
		ar := strings.Fields(l)
		if len(ar) < 3 {
			continue
		}

		if strings.ToLower(ar[1]) != "t" {
			continue
		}

		if _, ok := cache.funcs[ar[2]]; !ok {
			continue
		}

		adr, err := strconv.ParseUint(ar[0], 16, 64)
		if err != nil {
			continue
		}

		cache.ksym[adr] = ar[2]
		cache.addr = append(cache.addr, adr)
	}

	sort.Slice(cache.addr, func(i, j int) bool { return cache.addr[i] < cache.addr[j] })
}

// Ksym translates a kernel memory address into a kernel function name
func Ksym(addr uint64) string {

	cache.Lock()
	defer cache.Unlock()

	i := sort.Search(len(cache.addr), func(i int) bool { return cache.addr[i] > addr })
	if i == 0 || i == len(cache.addr) {
		return "0x" + strconv.FormatUint(addr, 16)
	}

	i = i - 1
	ofs := addr - cache.addr[i]
	return cache.ksym[cache.addr[i]] + ":0x" + strconv.FormatUint(ofs, 16)
}
