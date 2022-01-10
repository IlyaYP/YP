package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "0.0.1"

func main() {
	// допишите код
	var oldUsage = flag.Usage
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s version:%s\n", os.Args[0], version)
		oldUsage()
	}
	flag.Parse()
}
