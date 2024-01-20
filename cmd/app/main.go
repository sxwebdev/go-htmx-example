package main

import (
	"github.com/sxwebdev/go-htmx-example/internal/config"
	"github.com/sxwebdev/go-htmx-example/internal/server"
	"github.com/tkcrm/mx/cfg"
	"github.com/tkcrm/mx/launcher"
	"github.com/tkcrm/mx/logger"
	"github.com/tkcrm/mx/service"
	"github.com/tkcrm/mx/service/pingpong"
)

var (
	appName = "go-htmx-example"
	version = "local"
)

func main() {
	l := logger.NewExtended(
		logger.WithAppVersion(version),
		logger.WithAppName(appName),
	)

	conf := new(config.Config)
	if err := cfg.Load(conf, cfg.WithVersion(version)); err != nil {
		l.Fatalf("failed to load configuration: %s", err)
	}

	if err := run(l, conf); err != nil {
		l.Fatalf("failed to run server: %s", err)
	}
}

func run(l logger.ExtendedLogger, config *config.Config) error {
	// init launcher
	ln := launcher.New(
		launcher.WithVersion(version),
		launcher.WithName(appName),
		launcher.WithLogger(l),
		launcher.WithRunnerServicesSequence(launcher.RunnerServicesSequenceFifo),
		launcher.WithOpsConfig(config.Ops),
		launcher.WithAppStartStopLog(true),
	)

	// init static server
	staticServer := server.New(l, config)

	// register services
	ln.ServicesRunner().Register(
		service.New(service.WithService(pingpong.New(l))),
		service.New(service.WithService(staticServer)),
	)

	return ln.Run()
}
