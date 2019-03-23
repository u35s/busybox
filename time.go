package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func (app *applets) Applet_time(args []string) {
	var set = flag.NewFlagSet(args[0], flag.ExitOnError)
	var vflag = set.Bool("v", false, "verbose")
	set.Parse(args[1:])
	args = set.Args()

	var elapsed time.Duration
	var ru *syscall.Rusage

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	var retval int

	const default_format = "real\t%E\nuser\t%u\nsys\t%T"
	const posix_format = "real %e\nuser %U\nsys %S"
	const long_format = "\tCommand being timed: \"%C\"\n" +
		"\tUser time (seconds): %U\n" +
		"\tSystem time (seconds): %S\n" +
		"\tPercent of CPU this job got: %P\n" +
		"\tElapsed (wall clock) time (h:mm:ss or m:ss): %E\n" +
		"\tAverage shared text size (kbytes): %X\n" +
		"\tAverage unshared data size (kbytes): %D\n" +
		"\tAverage stack size (kbytes): %p\n" +
		"\tAverage total size (kbytes): %K\n" +
		"\tMaximum resident set size (kbytes): %M\n" +
		"\tAverage resident set size (kbytes): %t\n" +
		"\tMajor (requiring I/O) page faults: %F\n" +
		"\tMinor (reclaiming a frame) page faults: %R\n" +
		"\tVoluntary context switches: %w\n" +
		"\tInvoluntary context switches: %c\n" +
		"\tSwaps: %W\n" +
		"\tFile system inputs: %I\n" +
		"\tFile system outputs: %O\n" +
		"\tSocket messages sent: %s\n" +
		"\tSocket messages received: %r\n" +
		"\tSignals delivered: %k\n" +
		"\tPage size (bytes): %Z\n" +
		"\tExit status: %x\n"

	summarize := func(format, command string) {
		var vv_ms int64
		var cpu_ticks int64
		var pagesize = int64(os.Getpagesize())

		vv_ms = (ru.Utime.Sec+ru.Stime.Sec)*1000 +
			int64((ru.Utime.Usec+ru.Stime.Usec)/1000)

		cpu_ticks = vv_ms * 100 / 1000
		if cpu_ticks == 0 {
			cpu_ticks = 1
		}
		var key byte
		for {
			n := strings.Index(format, "%")
			if n > 0 {
				fmt.Printf("%s", format[:n])
				key = format[n+1]
				format = format[n+2:]
			} else {
				break
			}
			switch key {
			case 'C': /* The command that got timed.  */
				fmt.Printf("%s", command)
			case 'D': /* Average unshared data size.  */
				fmt.Printf("%d", (ru.Isrss+ru.Idrss)*pagesize/1024/cpu_ticks)
			case 'E': /* Average unshared data size.  */
				seconds := int64(elapsed.Seconds())
				if seconds >= 3600 {
					fmt.Printf("%dh %dm %ds", seconds/3600, (seconds%3600)/60, seconds%60)
				} else {
					fmt.Printf("%dm %.2fs", seconds/60, float64(elapsed.Nanoseconds())/1e9)
				}
			case 'F': /* Major page faults.  */
				fmt.Printf("%d", ru.Majflt)
			case 'I': /* Inputs.  */
				fmt.Printf("%d", ru.Inblock)
			case 'K': /* Average mem usage == data+stack+text.  */
				fmt.Printf("%d", (ru.Idrss+ru.Isrss+ru.Ixrss)*pagesize/cpu_ticks)
			case 'O': /* Maximum resident set size.  */
				fmt.Printf("%d", ru.Oublock)
			case 'P': /* Percent of CPU this job got.  */
				/* % cpu is (total cpu time)/(elapsed time).  */
				fmt.Printf("%d", vv_ms*100*1e6/int64(elapsed))
			case 'R': /* Minor page faults (reclaims).  */
				fmt.Printf("%d", ru.Minflt)
			case 'S': /* System time.  */
				fmt.Printf("%d.%d", ru.Stime.Sec, ru.Stime.Usec/10000)
			case 'T': /* System time.  */
				seconds := ru.Stime.Sec
				if seconds >= 3600 {
					fmt.Printf("%dh %dm %ds", seconds/3600, (seconds%3600)/60, seconds%60)
				} else {
					fmt.Printf("%dm %d.%ds", seconds/60, seconds%60, ru.Stime.Usec/1e7)
				}
			case 'U': /* User time.  */
				fmt.Printf("%d.%d", ru.Utime.Sec, ru.Utime.Usec/10000)
			case 'u': /* User time.  */
				seconds := ru.Utime.Sec
				if seconds >= 3600 {
					fmt.Printf("%dh %dm %ds", seconds/3600, (seconds%3600)/60, seconds%60)
				} else {
					fmt.Printf("%dm %d.%ds", seconds/60, seconds%60, ru.Utime.Usec/1e7)
				}
			case 'W': /* Times swapped out.  */
				fmt.Printf("%d", ru.Nswap)
			case 'X': /* Average shared text size.  */
				fmt.Printf("%d", ru.Ixrss*pagesize/cpu_ticks)
			case 'Z': /* Page size.  */
				fmt.Printf("%d", pagesize)
			case 'c': /* Involuntary context switches.  */
				fmt.Printf("%d", ru.Nivcsw)
			case 'e': /* Elapsed real time in seconds.  */
				fmt.Printf("%.2f", float64(elapsed.Nanoseconds())/1e9)
			case 'k': /* Signals delivered.  */
				fmt.Printf("%d", ru.Nsignals)
			case 'p': /* Average stack segment.  */
				fmt.Printf("%d", ru.Isrss*pagesize/cpu_ticks)
			case 'r': /* Incoming socket messages received.  */
				fmt.Printf("%d", ru.Msgrcv)
			case 's': /* Outgoing socket messages sent.  */
				fmt.Printf("%d", ru.Msgsnd)
			case 't': /* Average resident set size.  */
				fmt.Printf("%d", ru.Idrss*pagesize/cpu_ticks)
			case 'w': /* Voluntary context switches.  */
				fmt.Printf("%d", ru.Nvcsw)
			case 'x': /* Exit status.  */
				fmt.Printf("%d", cmd.ProcessState.ExitCode())
			}
		}
		fmt.Printf("\n")
	}

	var start = time.Now()
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		retval = 1
	} else {
		elapsed = time.Since(start)
		ru = cmd.ProcessState.SysUsage().(*syscall.Rusage)

		if *vflag {
			summarize(long_format, args[0])
		} else {
			summarize(posix_format, args[0])
		}
	}

	os.Exit(retval)
}
