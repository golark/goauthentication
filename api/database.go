package main

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v4"
)

func ConnectPgx(connStr string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	return conn, nil
}

func ConnectSQL(dsn string) (*sql.DB, error) {
	pool, err := sql.Open("postgress", dsn)
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	ctx := context.Background()
	err = pool.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
