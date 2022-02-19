package main

import (
	"crypto/tls"
	"github.com/go-redis/redis/v7"
)

func NewRedisClient(host string, pass string) (*redis.Client, error) {

	op := &redis.Options{Addr: host, Password: pass, TLSConfig: &tls.Config{MinVersion: tls.VersionTLS12}}

	client := redis.NewClient(op)

	err := client.Ping().Err()
	if err != nil {
		return nil, err
	}

	return client, nil
}
