package infrastructure

import (
	"database/sql"
	"devquest-server/config"
	"time"

	pgcommands "github.com/habx/pg-commands"
)

type Database interface {
	GetDB() *sql.DB
	GetDBTimeout() time.Duration
	CreateBackup(*config.Config) pgcommands.Result
}