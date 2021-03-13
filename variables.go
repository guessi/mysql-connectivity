package main

import "time"

var (
	mysql_user                 string
	mysql_pass                 string
	mysql_addr                 string
	mysql_dbname               string
	mysql_command              string
	mysql_timeout              = time.Millisecond * 1000
	mysql_readTimeout          = time.Millisecond * 1000
	mysql_writeTimeout         = time.Millisecond * 1000
	mysql_net                  = "tcp"
	mysql_allowNativePasswords = true
	mysql_dsn                  string
	check_interval             = time.Millisecond * 1000
)
