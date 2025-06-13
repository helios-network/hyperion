package main

import (
	"fmt"
	"os"

	cli "github.com/jawher/mow.cli"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/version"
)

var app = cli.App("hyperion", "Hyperion is a companion executable for orchestrating a Hyperion validator.")

var (
	envName        *string
	appLogLevel    *string
	svcWaitTimeout *string
)

func main() {
	readEnv()
	initGlobalOptions(
		&envName,
		&appLogLevel,
		&svcWaitTimeout,
	)

	app.Before = func() {
		log.DefaultLogger.SetLevel(logLevel(*appLogLevel))
	}

	app.Command("version", "Print the version information and exit.", versionCmd)
	app.Command("server", "Starts the server.", startServer)
	app.Command("cancel-all-pending-out-tx", "Cancels all pending outgoing txs.", cancelAllPendingOutTxCmd)

	_ = app.Run(os.Args)
}

func versionCmd(c *cli.Cmd) {
	c.Action = func() {
		fmt.Println(version.Version())
	}
}
