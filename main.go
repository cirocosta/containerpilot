package main // import "github.com/cirocosta/containerpilot"

import (
	"runtime"

	"github.com/cirocosta/containerpilot/core"
  "flag"

	log "github.com/Sirupsen/logrus"

	// Import backends so that they initialize
	_ "github.com/cirocosta/containerpilot/discovery/consul"
)

// Main executes the containerpilot CLI
func main() {
	// make sure we use only a single CPU so as not to cause
	// contention on the main application
	runtime.GOMAXPROCS(1)

	app, configErr := core.LoadApp()
	if configErr != nil {
		log.Fatal(configErr)
	}
	app.Run(flag.Args()) // Blocks forever
}
