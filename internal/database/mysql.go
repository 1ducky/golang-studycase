package database

import (
	"context"
	"database/sql"
	"fmt"
	"restApi/config"
	"time"

	_ "github.com/go-sql-driver/mysql" //initalize
)

func NewDatabase(config *config.DatabaseConfig) (*sql.DB, error) {
	fmt.Println(config)

	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.User, config.Password, config.Host, config.Port, config.DbName)
	fmt.Println(dsn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(15 * time.Minute)

	return db, nil
}
