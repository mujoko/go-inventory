package db

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"

	"go-inventory/config"
)

func New(conf *config.DbConf) (*sql.DB, error) {
	cfg := &mysql.Config{
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%v:%v", conf.Host, conf.Port),
		DBName:               conf.DbName,
		User:                 conf.Username,
		Passwd:               conf.Password,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	return sql.Open("mysql", cfg.FormatDSN())
}
