package main

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/dgraph-io/badger/v3"
	"github.com/paulbellamy/ratecounter"
	log "github.com/sirupsen/logrus"
)

func main() {
	argumentParsing(os.Args)

	// Open seencheck database
	tempDir, err := ioutil.TempDir(arguments.TempDir, "strainer-")
	if err != nil {
		log.Fatal(err)
	}
	seencheck := new(Seencheck)
	seencheck.SeenCount = new(ratecounter.Counter)
	seencheck.SeenDB, err = badger.Open(badger.DefaultOptions(tempDir))
	if err != nil {
		log.Fatal(err)
	}

	// Show statistics
	stats := new(Stats)
	stats.URIsPerSecond = ratecounter.NewRateCounter(1 * time.Second)
	stats.ParsedCounter = new(ratecounter.Counter)
	stats.DuplicateCounter = new(ratecounter.Counter)
	stats.ExcludedCounter = new(ratecounter.Counter)
	stats.SeedsListSize = new(ratecounter.Counter)
	stats.UniqueCounter = new(ratecounter.Counter)
	stats.HostsCount = new(ratecounter.Counter)
	stats.ParsingFailures = new(ratecounter.Counter)
	stats.StartTime = time.Now()
	stats.FilesCount = len(arguments.FrontierFiles)

	go stats.printLiveStats()

	// If the output file doesn't exist, create it, or append to the file
	fileName := "strainer_" + time.Now().Format("20060102150405") + ".txt"
	outFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// Process frontier file(s)
	for _, filePath := range arguments.FrontierFiles {
		stats.FileProcessingCount++
		stats.FilePath = filePath
		process(filePath, outFile, seencheck, stats)
	}

	os.RemoveAll(tempDir)

	time.Sleep(2 * time.Second)
}
