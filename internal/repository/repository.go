package repository

import (
	"context"
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/murtll/mcserver-beat/config"
	"github.com/murtll/mcserver-beat/internal/utils"
	"github.com/redis/go-redis/v9"
)

var client = redis.NewClient(config.RedisOptions)
var ctx = context.Background()

func Store(names []string, ttl time.Duration) error {
	key := utils.JSONTime(time.Now()).RoundHour().String()

	tmp := make([]interface{}, len(names))
	for i, v := range names {
		tmp[i] = v
	}
	err := client.SAdd(ctx, key, tmp...).Err()
	if err != nil {
		return err
	}
	err = client.Expire(ctx, key, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

func Load(count int) (map[int][]string, error) {
	stringKeys, err := client.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	keys := make([]int, 0)
	for _, stringK := range stringKeys {
		k, err := strconv.Atoi(stringK)
		if err != nil {
			log.Default().Printf("Was not able to parse Redis key '%s': %s. Skipping.", stringK, err)
			continue
		}
		keys = append(keys, k)
	}

	if count > len(keys) {
		count = len(keys)
	}

	sort.Ints(keys)
	keys = keys[len(keys) - count:]

	result := make(map[int][]string)
	for _, k := range keys {
		result[k], err = client.SInter(ctx, strconv.Itoa(k)).Result()
		if err != nil {
			log.Default().Printf("Was not able to get Redis key '%d': %s. Skipping.", k, err)
			delete(result, k)
			continue
		}
	}

	return result, nil
}