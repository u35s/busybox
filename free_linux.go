package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"syscall"
)

func (app *applets) Applet_free(args []string) {
	var unit_steps uint = 10

	var set = flag.NewFlagSet(args[0], flag.ExitOnError)
	set.Bool("b", false, "display memory in Bytes")
	set.Bool("k", false, "display memory in KB")
	set.Bool("m", false, "display memory in MB")
	set.Bool("g", false, "display memory in GB")
	//var sflag = set.Int("s", 0, "refresh the output every few seconds")
	//var hflag = set.Bool("h", false, "display memory in human readable units")
	set.Parse(args[1:])
	set.VisitAll(func(flag *flag.Flag) {
		switch {
		case flag.Name == "b" && flag.Value.String() == "true":
			unit_steps = 0
		case flag.Name == "k" && flag.Value.String() == "true":
			unit_steps = 10
		case flag.Name == "m" && flag.Value.String() == "true":
			unit_steps = 20
		case flag.Name == "g" && flag.Value.String() == "true":
			unit_steps = 30
		}
	})
	args = set.Args()

	var retval int

	var info syscall.Sysinfo_t
	var cached, cached_plus_free, available uint64
	var cached_kb, available_kb uint64
	var seen_available bool

	syscall.Sysinfo(&info)
	// parse_meminfo
	{
		var seen_cached_and_available int = 2
		mfile, err := os.Open("/proc/meminfo")
		if err != nil {
			fmt.Println(err)
			retval = 1
		} else {
			reader := bufio.NewReader(mfile)
			for {
				line, _, err := reader.ReadLine()
				if err == io.EOF {
					break
				} else if err != nil {
					retval = 1
					break
				}
				if n, err1 := fmt.Sscanf(string(line), "Cached: %d %*s\n", &cached_kb); err1 == nil && n == 1 {
					fmt.Println("Cached", n, err1, cached_kb)
					seen_cached_and_available--
					if seen_cached_and_available == 0 {
						break
					}
				}
				if n, err1 := fmt.Sscanf(string(line), "MemAvailable: %d %*s\n", &available_kb); err1 == nil && n == 1 {
					fmt.Println("Mem", n, err1, cached_kb)
					seen_cached_and_available--
					if seen_cached_and_available == 0 {
						break
					}
				}
			}
			seen_available = seen_cached_and_available == 0
		}
	}
	available = (available_kb * 1024)
	cached = (cached_kb * 1024)
	cached += info.Bufferram
	cached_plus_free = cached + info.Freeram

	fmt.Printf("       %12s%12s%12s%12s%12s%12s\nMem:   ",
		"total",
		"used",
		"free",
		"shared", "buff/cache", "available")

	scale := func(d uint64) uint64 {
		return d >> unit_steps
	}

	var FIELDS_6 = "%12d %11d %11d %11d %11d %11d\n"
	var FIELDS_3 = "%12d %11d %11d\n"

	fmt.Printf(FIELDS_6,
		scale(info.Totalram),                  //total
		scale(info.Totalram-cached_plus_free), //used
		scale(info.Freeram),                   //free
		scale(info.Sharedram),                 //shared
		scale(cached),                         //buff/cache
		scale(available),                      //available
	)
	if !seen_available {
	}
	fmt.Printf("Swap:  ")
	fmt.Printf(FIELDS_3,
		scale(info.Totalswap),               //total
		scale(info.Totalswap-info.Freeswap), //used
		scale(info.Freeswap),                //free
	)

	os.Exit(retval)
}
