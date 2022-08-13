package main

import (
	"context"
	"os"

	flagKong "github.com/alecthomas/kong"
	"github.com/mrmm/act-artifact-server/pkg"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Parse the command line arguments
	flagKong.Parse(&pkg.CLI)

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Parse log level provided by the cli
	lvl, err := log.ParseLevel(pkg.CLI.Log.Level)
	if err != nil {
		log.Fatalf("Failed to parse log level: %s", pkg.CLI.Log.Level)
	}
	// Only log the warning severity or above.
	log.SetLevel(lvl)
}

func main() {

	log.Info("Starting artifact-server with Token: ", pkg.CLI.Server.Token)
	ctx := context.Background()
	cancel := pkg.Serve(ctx)

	defer cancel()

}
