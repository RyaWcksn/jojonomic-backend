package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/app/handler"
	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/app/service"
	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/architecture/message"
	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/config"
	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/domain/broker"
	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/domain/price"
	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/domain/storage"
	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/pkgs/logger"
	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/server/middleware"
	"github.com/gorilla/mux"
)

func SRV(cfg config.Config, log logger.ILogger) {
	addr := cfg.Host
	var signalChan chan (os.Signal) = make(chan os.Signal, 1)

	priceImpl := price.NewPrice(cfg, log)
	storageImpl := storage.NewStorage(cfg, log)
	kafkaConn := message.NewKafkaProducer(cfg, log)
	brokerImpl := broker.NewMessageBroker(kafkaConn, log)
	serviceImpl := service.NewService(brokerImpl, log, storageImpl, priceImpl)
	handler := handler.NewHandler(serviceImpl, log)

	r := mux.NewRouter()
	r.Handle("/api/buyback", middleware.ErrHandler(handler.BuybackHandler)).Methods(http.MethodPost)
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
