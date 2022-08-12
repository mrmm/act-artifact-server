package main

import (
	"context"
	"flag"
	"os"

	"github.com/mrmm/act-artifact-server/pkg"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	ctx := context.Background()

	var host, port, artifactPath string
	flag.StringVar(&host, "h", "0.0.0.0", "Host to listen on")
	flag.StringVar(&port, "p", "1234", "Port to listen on")
	flag.StringVar(&artifactPath, "d", "/tmp", "Artifact path")
	flag.Parse()

	cancel := pkg.Serve(ctx, artifactPath, host, port)

	defer cancel()

}
