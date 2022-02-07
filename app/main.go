package main

import (
	"github.com/go-redis/redis/v7"
	"log"
	"os"
)

type application struct {
	redisClient *redis.Client
}

func main() {
	redisHost := os.Getenv("REDIS_HOST_ADDRESS")
	redisPass := os.Getenv("REDIS_PASSWORD")

	// step 1 - get redis client
	if len(redisHost) == 0 {
		log.Fatal("cant find REDIS_HOST_ADDRESS in environment variables, exciting...")
	}
	if len(redisPass) == 0 {
		log.Fatal("cant find REDIS_PASSWORD in environment variables, exciting...")
	}
	redisClient, err := NewRedisClient(redisHost, redisPass)
	if err!= nil {
		log.Fatalf("cant ping redis %v", err)
	}
	defer func(client *redis.Client) {
		err := client.Close()
		if err != nil {

		}
	}(redisClient)

	// step 2 - create application
	app := application{
		redisClient: redisClient,
	}

	// step 3 - setup and run router
	router := setupGinRouter(app)
	log.Fatal(router.Run(":8080"))
}
