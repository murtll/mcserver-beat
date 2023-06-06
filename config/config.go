package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func GetStrOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		return defaultValue
	}
}

func GetIntOrDefault(key string, defaultValue int) int {
	if stringValue, ok := os.LookupEnv(key); ok {
		value, err := strconv.Atoi(stringValue)
		if err != nil {
			log.Default().Printf("Error converting string env '%s' to int. Falling back to default.\n", key)
			return defaultValue
		}
		return value
	} else {
		return defaultValue
	}
}

var ListenAddr = GetStrOrDefault("LISTEN_ADDR", ":1620")
var ListenPath = GetStrOrDefault("LISTEN_PATH", "/graphinfo")
var HealthPath = GetStrOrDefault("HEALTH_PATH", "/_healthz")

var redisHost = GetStrOrDefault("REDIS_HOST", "localhost")
var redisPort = GetStrOrDefault("REDIS_PORT", "6379")
var RedisOptions = &redis.Options{
	Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
	Password: GetStrOrDefault("REDIS_PASSWORD", ""),
	DB: GetIntOrDefault("REDIS_DB", 0),
}

var EntryTTL = time.Duration(GetIntOrDefault("ENTRY_TTL_HOURS", 10) * int(time.Hour))
var EntryNumber = GetIntOrDefault("ENTRY_NUMBER", 10)

var PollingInterval = time.Duration(GetIntOrDefault("POLLING_INTERVAL", 5) * int(time.Second))
var PollingURL = GetStrOrDefault("POLLING_URL", "https://api.minetools.eu/query/play.mcbrawl.ru/25565")
var PollingSchema = GetStrOrDefault("POLLING_SCHEMA", "Players")

var Version = GetStrOrDefault("APP_VERSION", "0.1.0")