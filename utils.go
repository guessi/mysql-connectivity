package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func init() {
	// setup log format with timestamp
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	// it will be helpful with debug message
	log.SetLevel(log.DebugLevel)
}

func execute() {
	// generate data source name (DSN)
	config := &mysql.Config{
		User:                 mysql_user,
		Passwd:               mysql_pass,
		Addr:                 mysql_addr,
		DBName:               mysql_dbname,
		Timeout:              mysql_timeout,
		ReadTimeout:          mysql_readTimeout,
		WriteTimeout:         mysql_writeTimeout,
		Net:                  mysql_net,
		AllowNativePasswords: mysql_allowNativePasswords,
	}
	mysql_dsn = config.FormatDSN()

	// execute command for checking db connectivity
	for range time.Tick(check_interval) {
		log.Info("Opening database connection")
		db, err := sql.Open("mysql", mysql_dsn)
		if err != nil {
			log.Error(err.Error())
			return
		}
		defer db.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// basic ping test before accessing `mysql_dbname`
		if err := db.PingContext(ctx); err != nil {
			log.Error(err.Error())
			return
		}
		log.Info("Ping pass")

		// try to execute `mysql_command` once per second
		for range time.Tick(check_interval) {
			var result string

			row := db.QueryRowContext(ctx, mysql_command)
			if err := row.Scan(&result); err != nil {
				log.Error(err.Error())
				return
			} else {
				log.Debug(result)
			}
		}
	}
}
