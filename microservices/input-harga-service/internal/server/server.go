package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/app/handler"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/app/service"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/architecture/message"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/config"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/domain/broker"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/pkgs/logger"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/server/middleware"
	"github.com/gorilla/mux"
)

func SRV(cfg config.Config, log logger.ILogger) {
	addr := cfg.Host
	var signalChan chan (os.Signal) = make(chan os.Signal, 1)

	kafkaConn := message.NewKafkaProducer(cfg, log)
	brokerImpl := broker.NewMessageBroker(kafkaConn, log)
	serviceImpl := service.NewService(brokerImpl, log)
	handler := handler.NewHandler(serviceImpl, log)

	r := mux.NewRouter()
	r.Handle("/api/input-harga/", middleware.ErrHandler(handler.InputHargaHandler))
	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	go func() {
		srv.ListenAndServe()
		log.Infof("HTTP server started %v", addr)
	}()

	sig := <-signalChan
	log.Infof("%s signal caught", sig)

	err := kafkaConn.Close()
	if err != nil {
		log.Errorf("Error in closing KAFKA connection. Err : %+v", err.Error())
	}

}
