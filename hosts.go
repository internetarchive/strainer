package main

import "github.com/paulbellamy/ratecounter"

var hosts []Host

type Host struct {
	Value   string
	Counter *ratecounter.Counter
}

func isHostExcluded(targetHost string, stats *Stats) bool {
	var found = false

	for _, host := range hosts {
		if host.Value == targetHost {
			found = true

			if host.Counter.Value() >= arguments.MaxHostOccurence {
				host.Counter.Incr(1)
				stats.ExcludedCounter.Incr(1)
				return true
			}

			host.Counter.Incr(1)
		}
	}

	if !found {
		newHost := Host{
			Value:   targetHost,
			Counter: new(ratecounter.Counter),
		}

		stats.HostsCount.Incr(1)
		newHost.Counter.Incr(1)

		hosts = append(hosts, newHost)
	}

	return false
}
