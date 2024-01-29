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
	// JwtAuthTokenSecretKey    string `validate:"required" secret:"true"`
	// JwtAuthTokenDuration     int    `validate:"required" default:"15" usage:"token expiration in minutes"`
	// JwtRefreshTokenSecretKey string `validate:"required" secret:"true"`
	// JwtRefreshTokenDuration  int    `validate:"required" default:"10080" usage:"token expiration in minutes. default 1 week"`
}
