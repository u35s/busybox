package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"syscall"
	"time"

	"github.com/u35s/busybox/libbb"
)

func (app *applets) Applet_stat(args []string) {
	var set = flag.NewFlagSet("stat [OPTION]... FILE...", flag.ExitOnError)
	var tflag = set.Bool("t", false, "print the information in terse form")
	/*ar fflag = set.Bool("f", false, "display file system status instead of file status")
	var cflag = set.String("c", "",
		"use the specified FORMAT instead of the default; output a newline after each use of FORMAT")
	*/
	set.Parse(args[1:])
	args = set.Args()

	var retval int

	getUserName := func(uid uint32) string {
		u, err := user.LookupId(fmt.Sprintf("%v", uid))
		if err != nil {
			return "UNKNOWN"
		}
		return u.Username
	}
	getGroupName := func(gid uint32) string {
		g, err := user.LookupGroupId(fmt.Sprintf("%v", gid))
		if err != nil {
			return "UNKNOWN"
		}
		return g.Name
	}
	humanTime := func(t syscall.Timespec) string {
		return time.Unix(t.Sec, t.Nsec).Format("2006-01-02 15:04:05.999999999 -0700")
	}
	fileType := func(m os.FileMode) string {
		switch {
		case m.IsRegular():
			return "regular file"
		case m.IsDir():
			return "directory"
		case m&syscall.S_IFBLK > 0:
			return "block special file"
		case m&syscall.S_IFCHR > 0:
			return "character special file"
		case m&syscall.S_IFIFO > 0:
			return "fifo"
		case m&syscall.S_IFLNK > 0:
			return "symbolic link"
		case m&syscall.S_IFSOCK > 0:
			return "socket"
		}
		return "wrird file"
	}

	doStat := func(info os.FileInfo) {
		sysStat := &libbb.Stat_t{info.Sys().(*syscall.Stat_t)}
		if *tflag {
			fmt.Printf("%s %d %d %x %d %d %x %d %d %x %x %d %d %d %d\n",
				info.Name(),
				info.Size(), sysStat.Blocks, sysStat.Mode, sysStat.Uid, sysStat.Gid,
				sysStat.Dev, sysStat.Ino, sysStat.Nlink, sysStat.Rdev, sysStat.Rdev,
				sysStat.AccessTime().Sec,
				sysStat.ModifyTime().Sec,
				sysStat.ChangeTime().Sec,
				sysStat.Blksize,
			)
		} else {
			var (
				linkname string
			)
			if info.Mode()&os.ModeSymlink > 0 {
				linkname, _ = os.Readlink(info.Name())
				fmt.Printf("  File: '%s' -> '%s'\n", info.Name(), linkname)
			} else {
				fmt.Printf("  File: '%s'\n", info.Name())
			}

			fmt.Printf("  Size: %-10d\tBlocks: %-10d IO Block: %-6d %s\n"+
				"Device: %xh/%dd\tInode: %-10d  Links: %-5d",
				sysStat.Size, sysStat.Blocks, sysStat.Blksize, fileType(info.Mode()),
				sysStat.Dev, sysStat.Dev, sysStat.Ino, sysStat.Nlink,
			)

			if info.Mode()&os.ModeDevice > 0 || info.Mode()&os.ModeCharDevice > 0 {
				fmt.Printf(" Device type: %x,%x\n",
					sysStat.Rdev, sysStat.Rdev)
			} else {
				fmt.Printf("\n")
			}

			fmt.Printf("Access: (%04o/%10.10s)  Uid: (%5d/%8s)   Gid: (%5d/%8s)\n",
				info.Mode()&(syscall.S_ISUID|syscall.S_ISGID|syscall.S_ISVTX|syscall.S_IRWXU|syscall.S_IRWXG|syscall.S_IRWXO),
				info.Mode(), sysStat.Uid, getUserName(sysStat.Uid), sysStat.Gid, getGroupName(sysStat.Gid),
			)

			fmt.Printf("Access: %s\n", humanTime(sysStat.AccessTime()))
			fmt.Printf("Modify: %s\n", humanTime(sysStat.ModifyTime()))
			fmt.Printf("Change: %s\n", humanTime(sysStat.ChangeTime()))
		}
	}

	for i := range args {
		file, err := os.Open(args[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			retval = 1
		} else {
			info, _ := file.Stat()
			doStat(info)
			file.Close()
		}
	}
	os.Exit(retval)
}
