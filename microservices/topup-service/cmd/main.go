package main

import (
	"github.com/RyaWcksn/jojonomic-backend/topup-service/internal/config"
	"github.com/RyaWcksn/jojonomic-backend/topup-service/internal/logger"
	"github.com/RyaWcksn/jojonomic-backend/topup-service/internal/server"
)

func main() {
	cfg := config.InitCfg()
	log := logger.Init(cfg.Service, cfg.Env, cfg.LogLevel)
	server.SRV(*cfg, log)
}
