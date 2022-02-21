package main

import (
	"flag"
	"github.com/go-redis/redis/v7"
	"log"
	"os"
	"strconv"
)

type application struct {
	redisClient *redis.Client
	hmacSecret  string
	infoLog     *log.Logger
}

func main() {
	redisHost := os.Getenv("REDIS_HOST_ADDRESS")
	redisPass := os.Getenv("REDIS_PASSWORD")
	hmacPassword := os.Getenv("HMAC_PASSWORD")

	var port int
	flag.IntVar(&port, "port", 5023, "port to listen on")

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
	infoLog := log.New(os.Stdout, "INFO:", log.Ldate)
	app := application{
		redisClient: redisClient,
		hmacSecret:  hmacPassword,
		infoLog:     infoLog,
	}

	// step 3 - setup and run router
	router := setupGinRouter(app)
	log.Fatal(router.Run(":" + strconv.Itoa(port)))
}
