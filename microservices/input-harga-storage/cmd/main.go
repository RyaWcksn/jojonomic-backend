package main

import (
	"fmt"

	"github.com/RyaWcksn/jojonomic-backend/input-harga-storage/internal/config"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-storage/internal/logger"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-storage/server"
)

func main() {
	cfg := config.InitCfg()
	fmt.Println(&cfg)
	log := logger.Init(cfg.Service, cfg.Env, cfg.LogLevel)
	server.SRV(*cfg, log)
}
