package main

import (
	"flag"
	"fmt"
	"runtime"
)

var (
	Version   = "0.0.0"
	CommitSha = "xxxx"
	BuildTime = "0000-00-00T00:00:00+00:00"
)

func main() {
	var (
		all        bool
		help       bool
		ignore     int
		getVersion bool
		noOfCpus   int
	)

	flag.BoolVar(&all, "all", true, "print the number of installed processors")
	flag.BoolVar(&help, "help", false, "display this help and exit")
	flag.BoolVar(&getVersion, "version", false, "Get current version")
	flag.IntVar(&ignore, "ignore", 0, "if possible, exclude N processing units")

	flag.Parse()

	if getVersion {
		fmt.Printf("go-nproc v%s (sha: %s) (BuildTime: %s)\n\nWritten by Pratik (github: pratikbin)\n",
			Version, CommitSha, BuildTime)
		return
	}
	if help {
		fmt.Println("go-nproc Usage")
		flag.PrintDefaults()
		return
	}

	noOfCpus = runtime.NumCPU()
	if ignore > 0 {
		if ignore >= noOfCpus {
			fmt.Println(1)
		} else {
			fmt.Println(noOfCpus - ignore)
		}
	} else if ignore < 0 {
		fmt.Printf("go-nproc: invalid number: '%v'", ignore)
	} else {
		fmt.Println(noOfCpus)
	}
}
