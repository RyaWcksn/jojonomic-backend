package main

import (
	"github.com/RyaWcksn/jojonomic-backend/check-harga-service/internal/config"
	"github.com/RyaWcksn/jojonomic-backend/check-harga-service/internal/logger"
	"github.com/RyaWcksn/jojonomic-backend/check-harga-service/internal/server"
)

func main() {
	cfg := config.InitCfg()
	log := logger.Init(cfg.Service, cfg.Env, cfg.LogLevel)
	server.SRV(*cfg, log)
}
