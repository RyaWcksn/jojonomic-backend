package main

import (
	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/config"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/pkgs/logger"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/server"
)

func main() {
	cfg := config.InitCfg()
	log := logger.Init(cfg.Service, cfg.Env, cfg.LogLevel)
	server.SRV(*cfg, log)
}
