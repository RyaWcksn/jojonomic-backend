package main

import (
	"github.com/RyaWcksn/jojonomic-backend/buyback-storage/internal/config"
	"github.com/RyaWcksn/jojonomic-backend/buyback-storage/internal/pkgs/logger"
	"github.com/RyaWcksn/jojonomic-backend/buyback-storage/internal/server"
)

func main() {
	cfg := config.InitCfg()
	log := logger.Init(cfg.Service, cfg.Env, cfg.LogLevel)
	server.SRV(*cfg, log)
}
