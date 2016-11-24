// Command line interface to run Cross Post Detector
package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/janosgyerik/rssreader"
)

const defaultConfigFile = "feeds.yml"

func exit() {
	flag.Usage()
	os.Exit(1)
}

type Params struct {
	configfile string
}

func parseArgs() Params {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options]\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	configfilePtr := flag.String("config", defaultConfigFile, "path to configuration file")
	flag.Parse()

	if len(flag.Args()) != 0 {
		exit()
	}

	return Params{*configfilePtr}
}

func main() {
	params := parseArgs()

	if err := rssreader.RunForever(params.configfile); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
