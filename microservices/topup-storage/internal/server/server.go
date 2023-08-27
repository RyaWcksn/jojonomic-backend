package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/app/service"
	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/architecture/database"
	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/architecture/message"
	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/config"
	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/domain/broker"
	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/domain/storage"
	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/domain/transaction"
	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/errors"
	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/logger"
)

func SRV(cfg config.Config, log logger.ILogger) {
	kafkaConn := message.NewKafkaReader(cfg, log)
	defer kafkaConn.Close()
	db := database.NewDatabaseConnection(cfg, log)
	dbConn := db.DBConnect()
	defer dbConn.Close()

	transactionImpl := transaction.NewTransaction(dbConn, log)
	storageImpl := storage.NewStorage(cfg, log)
	brokerImpl := broker.NewBrokerImpl(kafkaConn, log)
	serviceImpl := service.NewService(brokerImpl, log, storageImpl, transactionImpl)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle OS signals to gracefully shutdown
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
		<-signals
		cancel()
	}()

	serviceImpl.Consume(ctx)

	defer func() {
		if err := recover(); err != nil {
			xerr := errors.ErrorForm{
				IsError: true,
				ReffId:  "",
				Message: "Panic",
			}
			log.Errorf("ERR := %v", xerr)
		}
	}()
}
