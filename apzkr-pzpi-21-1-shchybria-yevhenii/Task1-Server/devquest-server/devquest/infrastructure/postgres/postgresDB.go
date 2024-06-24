package postgres

import (
	"database/sql"
	"devquest-server/config"
	"fmt"
	"os"
	"sync"
	"time"

	pgcommands "github.com/habx/pg-commands"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresDB struct {
	Db *sql.DB
}

var (
	once       sync.Once
	dbInstance *PostgresDB
	dbError    error
)

func NewPostgresDB(conf *config.Config) (*PostgresDB, error) {
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s timezone=%s connect_timeout=%d",
			conf.Database.Host,
			conf.Database.Port,
			conf.Database.User,
			conf.Database.Password,
			conf.Database.DBName,
			conf.Database.SSLMode,
			conf.Database.TimeZone,
			conf.Database.ConnectTimeout,
		)

		db, err := sql.Open("pgx", dsn)
		if err != nil {
			dbError = err
			return
		}

		if err = db.Ping(); err != nil {
			dbError = err
			return
		}

		dbInstance = &PostgresDB{Db: db}
	})

	return dbInstance, dbError
}

func (p *PostgresDB) GetDB() *sql.DB {
	return p.Db
}

func (p *PostgresDB) GetDBTimeout() time.Duration {
	return time.Second * 3
}

func (p *PostgresDB) CreateBackup(conf *config.Config) pgcommands.Result {
	dump, _ := pgcommands.NewDump(&pgcommands.Postgres{
		Host: conf.Database.Host,
		Port: conf.Database.Port,
		DB: conf.Database.DBName,
		Username: conf.Database.User,
		Password: conf.Database.Password,
		EnvPassword: conf.Database.Password,
	})

	currentDir, _ := os.Getwd()

	dump.Options = []string{}
	dump.SetupFormat("t")
	dump.SetFileName(fmt.Sprintf(`%v_%v.sql.tar`, dump.DB, time.Now().Unix()))
	dump.SetPath(fmt.Sprintf("%s\\backups\\", currentDir))
	
	dumpExec := dump.Exec(pgcommands.ExecOptions{StreamPrint: true, StreamDestination: os.Stdout})
	return dumpExec
}