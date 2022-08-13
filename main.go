package main

import (
	"context"
	"os"

	"github.com/mrmm/act-artifact-server/pkg"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Parse the command line arguments
	pkg.ArgsApp.Parse(os.Args[1:])

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {

	log.Debug("Starting artifact-server with Token: ", *pkg.ArgsAuthToken)
	ctx := context.Background()
	cancel := pkg.Serve(ctx)

	defer cancel()

}
