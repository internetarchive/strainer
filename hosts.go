package main

import "github.com/paulbellamy/ratecounter"

var hosts []Host

type Host struct {
	Value   string
	Counter *ratecounter.Counter
}

func isHostExcluded(targetHost string, stats *Stats) bool {
	var found = false

	// Check if the host is simply excluded
	for _, excludedHost := range arguments.ExcludedHosts {
		if excludedHost == targetHost {
			stats.ExcludedCounter.Incr(1)

			for _, host := range hosts {
				if host.Value == targetHost {
					found = true
				}
			}

			if !found {
				stats.HostsCount.Incr(1)
			}

			return true
		}
	}

	for _, host := range hosts {
		if host.Value == targetHost {
			found = true

			// Check if the number of occurence of the host reached the --max-host-occurence limit
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
