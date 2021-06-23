package main

import (
	"bufio"
	"compress/gzip"
	"net/url"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/zeebo/xxh3"
)

func process(path string, outFile *os.File, seencheck *Seencheck, stats *Stats) {
	var scanner *bufio.Scanner

	// Open frontier file
	frontier, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer frontier.Close()

	// If the file ends with .gz, we open it has a GZIP file
	if strings.HasSuffix(path, ".gz") {
		reader, err := gzip.NewReader(frontier)
		if err != nil {
			log.Fatal(err)
		}
		defer reader.Close()
		scanner = bufio.NewScanner(reader)
	} else {
		scanner = bufio.NewScanner(frontier)
	}

	for scanner.Scan() {
		stats.ParsedCounter.Incr(1)
		stats.URIsPerSecond.Incr(1)
		rawURL := strings.Split(scanner.Text(), " ")
		if strings.Compare(rawURL[0], "F+") == 0 {
			// Parse URL
			URL, err := url.Parse(rawURL[1])
			if err != nil {
				stats.ParsingFailures.Incr(1)
				continue
			}

			// Generate URL hash
			hash := strconv.FormatUint(xxh3.HashString(URL.String()), 10)

			// Check if we already saw the URL, is yes then we skip it, else we add it to the seed list
			found, _, err := seencheck.IsSeen(hash)
			if err != nil {
				log.Fatal(err)
			}
			if !found {
				// Check host to see if we should exclude the URL
				if arguments.MaxHostOccurrence != -1 || len(arguments.ExcludedHosts) > 0 {
					if isHostExcluded(URL.Host, stats) {
						stats.UniqueCounter.Incr(1)
						seencheck.Seen(hash, URL.Host)
						continue
					}
				}

				if _, err = outFile.WriteString(URL.String() + "\n"); err != nil {
					panic(err)
				}
				stats.UniqueCounter.Incr(1)
				stats.SeedsListSize.Incr(1)
				seencheck.Seen(hash, URL.Host)
			} else {
				stats.DuplicateCounter.Incr(1)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
