[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/internetarchive/strainer.svg)](https://github.com/internetarchive/strainer)
[![Go Report Card](https://goreportcard.com/badge/github.com/internetarchive/strainer)](https://goreportcard.com/report/github.com/internetarchive/strainer)
[![GitHub license](https://img.shields.io/github/license/internetarchive/strainer.svg)](https://github.com/internetarchive/strainer/blob/master/LICENSE)

# Strainer

Parse Heritrix frontiers, build seeds list, exclude hosts, crawl better.

## What's strainer?

Strainer is a tool built to manipulate Heritrix frontier files.
It takes one or multiple (GZIP or not) frontier files and build new seeds list based on them.

It can exclude specific hosts, limit the number of occurrence of hosts in the final seeds list, and give you interesting statistics about your frontiers.

It's still a WIP, but it's usable right now.

## Usage

```
usage: strainer [-h|--help] -f|--file "<value>" [-f|--file "<value>" ...]
                [-m|--max-host-occurrence <integer>] [-e|--excluded-hosts
                "<value>" [-e|--excluded-hosts "<value>" ...]] [--temp-dir
                "<value>"]

                manipulate Heritrix frontier files

Arguments:

  -h  --help                 Print help information
  -f  --file                 Frontier file(s) to process, can be .gz files.
  -m  --max-host-occurrence  Max number of a occurrence of a given host to
                             accept in the final seed list. If an host is
                             parsed more than X times, new occurrences of that
                             host past that limit will be excluded. -1 value
                             means no limit. Default: -1
  -e  --excluded-hosts       Specific hosts to exclude from the final seed
                             list.
      --temp-dir             Temporary directory to use for the key/value
                             database. Default: /tmp
```
