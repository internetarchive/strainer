package main

import (
	"os"

	"github.com/akamensky/argparse"
	log "github.com/sirupsen/logrus"
)

var arguments struct {
	MaxHostOccurence int64
	FrontierFiles    []string
	ExcludedHosts    []string
	ShowStats        bool
}

func argumentParsing(args []string) {
	// Create new parser object
	parser := argparse.NewParser("strainer", "manipulate H3 frontier files")

	frontierFiles := parser.StringList("f", "file", &argparse.Options{
		Required: true,
		Help:     "Frontier file(s) to process, can be .gz files."})

	maxHostOccurence := parser.Int("m", "max-host-occurence", &argparse.Options{
		Required: false,
		Default:  -1,
		Help:     "Max number of a occurence of a given host to accept in the final seeds list. If an host is parsed more than X times, new occurences of that host past that limit will be excluded. -1 value means no limit.",
	})

	excludedHosts := parser.StringList("e", "excluded-hosts", &argparse.Options{
		Required: false,
		Help:     "Specific hosts to exclude from the final result.",
	})

	// Parse input
	err := parser.Parse(args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		log.Error(parser.Usage(err))
		os.Exit(0)
	}

	// Finally save the collected flags
	arguments.FrontierFiles = *frontierFiles
	arguments.MaxHostOccurence = int64(*maxHostOccurence)
	arguments.ExcludedHosts = *excludedHosts
}
