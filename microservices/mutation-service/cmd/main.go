package main

import (
	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/config"
	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/pkgs/logger"
	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/server"
)

func main() {
	cfg := config.InitCfg()
	log := logger.Init(cfg.Service, cfg.Env, cfg.LogLevel)
	server.SRV(*cfg, log)
}
