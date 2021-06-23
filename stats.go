package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gosuri/uilive"
	"github.com/gosuri/uitable"
	"github.com/paulbellamy/ratecounter"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Stats struct {
	URIsPerSecond       *ratecounter.RateCounter
	ParsedCounter       *ratecounter.Counter
	DuplicateCounter    *ratecounter.Counter
	UniqueCounter       *ratecounter.Counter
	HostsCount          *ratecounter.Counter
	ExcludedCounter     *ratecounter.Counter
	SeedsListSize       *ratecounter.Counter
	ParsingFailures     *ratecounter.Counter
	StartTime           time.Time
	FilePath            string
	FileProcessingCount int
	FilesCount          int
}

func (s *Stats) printLiveStats() {
	var stats *uitable.Table
	var m runtime.MemStats

	p := message.NewPrinter(language.English)

	writer := uilive.New()
	writer.Start()

	for {
		runtime.ReadMemStats(&m)

		stats = uitable.New()
		stats.MaxColWidth = 80
		stats.Wrap = true

		stats.AddRow("", "")
		stats.AddRow("  - File:", fmt.Sprintf("%s [%d/%d]", s.FilePath, s.FileProcessingCount, s.FilesCount))
		stats.AddRow("  - Speed:", p.Sprintf("%d URI/s", s.URIsPerSecond.Rate()))
		stats.AddRow("  - Parsed:", p.Sprintf("%d URIs", s.ParsedCounter.Value()))
		stats.AddRow("  - Unique:", p.Sprintf("%d URIs", s.UniqueCounter.Value()))
		stats.AddRow("  - Duplicate:", p.Sprintf("%d URIs", s.DuplicateCounter.Value()))

		if arguments.MaxHostOccurrence != -1 || len(arguments.ExcludedHosts) > 0 {
			stats.AddRow("  - Excluded:", p.Sprintf("%d URIs", s.ExcludedCounter.Value()))
			stats.AddRow("  - Unique hosts:", p.Sprintf("%d", s.HostsCount.Value()))
		}

		stats.AddRow("", "")
		stats.AddRow("  - Final seeds list size:", p.Sprintf("%d URIs", s.SeedsListSize.Value()))
		stats.AddRow("  - Elapsed time:", time.Since(s.StartTime))

		fmt.Fprintln(writer, stats.String())
		writer.Flush()
		time.Sleep(time.Second / 4)
	}
}
