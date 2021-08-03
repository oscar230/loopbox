package redis

import (
	"loopbox/utils/logger"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

func Client() *redis.Client {
	rHost := os.Getenv("REDIS_HOST")
	rPort := os.Getenv("REDIS_PORT")
	rPass := os.Getenv("REDIS_PASS")
	rDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	client := redis.NewClient(&redis.Options{
		Addr:     rHost + ":" + rPort,
		Password: rPass,
		DB:       rDB,
	})
	if _, rErr := client.Ping().Result(); rErr != nil {
		logger.Error("[redis] Cannot ping: " + rErr.Error())
	}
	return client
}
