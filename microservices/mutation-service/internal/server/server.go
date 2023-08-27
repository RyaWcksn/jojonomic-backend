package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/app/handler"
	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/app/service"
	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/architecture/database"
	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/config"
	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/domain/storage"
	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/pkgs/logger"
	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/server/middleware"
	"github.com/gorilla/mux"
)

func SRV(cfg config.Config, log logger.ILogger) {
	addr := cfg.Host
	var signalChan chan (os.Signal) = make(chan os.Signal, 1)

	db := database.NewDatabaseConnection(cfg, log)
	dbConn := db.DBConnect()

	storageImpl := storage.NewStorage(dbConn, log)
	serviceImpl := service.NewService(storageImpl, log)
	handler := handler.NewHandler(serviceImpl, log)

	r := mux.NewRouter()
	r.Handle("/api/mutasi", middleware.ErrHandler(handler.FetchMutationHandler)).Methods(http.MethodPost)
	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Infof("HTTP server started %v:%v", addr, cfg.Port)
	go func() {
		srv.ListenAndServe()
	}()

	sig := <-signalChan
	log.Infof("%s signal caught", sig)

}
