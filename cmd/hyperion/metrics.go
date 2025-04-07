package main

import (
	"os"
	"time"

	"github.com/Helios-Chain-Labs/metrics"
	cli "github.com/jawher/mow.cli"
	"github.com/xlab/closer"
	log "github.com/xlab/suplog"
)

func initMetrics(c *cli.Cmd) {
	var (
		statsdAgent    *string
		statsdPrefix   *string
		statsdAddr     *string
		statsdStuckDur *string
		statsdMocking  *string
		statsdDisabled *string
	)

	initStatsdOptions(
		c,
		&statsdAgent,
		&statsdPrefix,
		&statsdAddr,
		&statsdStuckDur,
		&statsdMocking,
		&statsdDisabled,
	)

	if !toBool(*statsdDisabled) {
		go func() {
			for {
				hostname, _ := os.Hostname()
				err := metrics.Init(*statsdAddr, checkStatsdPrefix(*statsdPrefix), &metrics.StatterConfig{
					Agent:                *statsdAgent,
					EnvName:              *envName,
					HostName:             hostname,
					StuckFunctionTimeout: duration(*statsdStuckDur, 30*time.Minute),
					MockingEnabled:       toBool(*statsdMocking),
				})
				if err != nil {
					log.WithError(err).Warningln("metrics init failed, will retry in 1 min")
					time.Sleep(time.Minute)
					continue
				}
				break
			}
			closer.Bind(func() {
				metrics.Close()
			})
		}()
	}
}
