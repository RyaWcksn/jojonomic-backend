package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/RyaWcksn/jojonomic-backend/input-harga-storage/internal/config"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-storage/internal/logger"
	_ "github.com/lib/pq"
)

type Connection struct {
	L   logger.ILogger
	cfg config.Config
}

func NewDatabaseConnection(M config.Config, l logger.ILogger) *Connection {
	return &Connection{
		cfg: M,
		L:   l,
	}
}

func (db *Connection) DBConnect() *sql.DB {
	addr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		db.cfg.SQLAddr, db.cfg.SQLPort, db.cfg.SQLUser, db.cfg.SQLPassword, db.cfg.SQLDatabase)

	dbConn, err := sql.Open("postgres", addr)
	if err != nil {
		db.L.Errorf("Error while connecting to DB... := %v", err)
		return nil
	}
	for dbConn.Ping() != nil {
		db.L.Info("Attempting connect to DB...")
		time.Sleep(5 * time.Second)
	}
	dbConn.SetMaxIdleConns(db.cfg.MaxConnection)
	dbConn.SetMaxOpenConns(db.cfg.MaxOpen)
	return dbConn
}
