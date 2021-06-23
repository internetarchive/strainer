package main

import (
	"os"

	"github.com/akamensky/argparse"
	log "github.com/sirupsen/logrus"
)

var arguments struct {
	MaxHostOccurrence int64
	TempDir           string
	FrontierFiles     []string
	ExcludedHosts     []string
	ShowStats         bool
}

func argumentParsing(args []string) {
	parser := argparse.NewParser("strainer", "manipulate Heritrix frontier files")

	frontierFiles := parser.StringList("f", "file", &argparse.Options{
		Required: true,
		Help:     "Frontier file(s) to process, can be .gz files."})

	maxHostOccurrence := parser.Int("m", "max-host-occurrence", &argparse.Options{
		Required: false,
		Default:  -1,
		Help:     "Max number of a occurrence of a given host to accept in the final seed list. If an host is parsed more than X times, new occurrences of that host past that limit will be excluded. -1 value means no limit",
	})

	excludedHosts := parser.StringList("e", "excluded-hosts", &argparse.Options{
		Required: false,
		Help:     "Specific hosts to exclude from the final seed list.",
	})

	tempDir := parser.String("", "temp-dir", &argparse.Options{
		Required: false,
		Default:  "/tmp",
		Help:     "Temporary directory to use for the key/value database",
	})

	// Parse input
	err := parser.Parse(args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		log.Error(parser.Usage(err))
		os.Exit(0)
	}

	// Test if all specified frontier files exist
	for _, path := range *frontierFiles {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			log.Fatalf("Specified frontier file %s do not exist!", path)
		}
	}

	// Finally save the collected flags
	arguments.FrontierFiles = *frontierFiles
	arguments.MaxHostOccurrence = int64(*maxHostOccurrence)
	arguments.ExcludedHosts = *excludedHosts
	arguments.TempDir = *tempDir
}
