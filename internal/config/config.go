package config

import (
	"github.com/tkcrm/mx/logger"
	"github.com/tkcrm/mx/ops"
)

// Config ...
type Config struct {
	EnvCI            string `validate:"required" env:"ENV_CI" example:"dev"`
	Log              logger.Config
	Ops              ops.Config
	StaticServerAddr string `validate:"required,hostname_port" default:":8080"`
}
