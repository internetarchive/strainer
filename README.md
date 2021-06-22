# Strainer

Parse Heritrix frontiers, build seeds list, exclude hosts, crawl better.

## What's strainer?

Strainer is a tool built to manipulate Heritrix frontier files.
It takes one or multiple (GZIP or not) frontier files and build new seeds list based on them.

It can exclude specific hosts, limit the number of occurence of hosts in the final seeds list, and give you interesting statistics about your frontiers.

It's still a WIP, but it's usable right now.

## Usage

```
usage: strainer [-h|--help] -f|--file "<value>" [-f|--file "<value>" ...]
                [-m|--max-host-occurence <integer>] [-e|--excluded-hosts
                "<value>" [-e|--excluded-hosts "<value>" ...]]

                manipulate H3 frontier files

Arguments:

  -h  --help                Print help information
  -f  --file                Frontier file(s) to process, can be .gz files.
  -m  --max-host-occurence  Max number of a occurence of a given host to accept
                            in the final seed list. If an host is parsed more
                            than X times, new occurences of that host past that
                            limit will be excluded. -1 value means no limit.
                            Default: -1
  -e  --excluded-hosts      Specific hosts to exclude from the final seed list.
```