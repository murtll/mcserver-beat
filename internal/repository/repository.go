package repository

import (
	"context"
	"sort"
	"strconv"
	"time"

	"github.com/murtll/mcserver-beat/config"
	"github.com/murtll/mcserver-beat/internal/utils"
	"github.com/redis/go-redis/v9"
)

var client = redis.NewClient(config.RedisOptions)
var ctx = context.Background()

func Store(count int, ttl time.Duration) error {
	key := utils.JSONTime(time.Now()).RoundHour().String()
	stringCurrentCount, err := client.Get(ctx, key).Result()

	if err != nil {
		if err == redis.Nil {
			err = client.Set(ctx, key, count, ttl).Err()
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	currentCount, err := strconv.Atoi(stringCurrentCount)
	if err != nil {
		return err
	}

	if currentCount < count {
		ttl, err := client.TTL(ctx, key).Result()
		if err != nil {
			return err
		}
		err = client.Set(ctx, key, count, ttl).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func Load(count int) (map[int]int, error) {
	stringKeys, err := client.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	if count > len(stringKeys) {
		count = len(stringKeys)
	}

	keys := make([]int, 0)
	for _, stringK := range stringKeys {
		k, err := strconv.Atoi(stringK)
		if err != nil {
			return nil, err
		}
		keys = append(keys, k)
	}

	sort.Ints(keys)
	keys = keys[len(keys) - count:]

	result := make(map[int]int)
	for _, k := range keys {
		value, err := client.Get(ctx, strconv.Itoa(k)).Result()
		if err != nil {
			return nil, err
		}
		result[k], err = strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}