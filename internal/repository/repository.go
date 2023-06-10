package repository

import (
	"context"
	"log"
	"time"

	"github.com/murtll/mcserver-beat/config"
	"github.com/murtll/mcserver-beat/internal/utils"
	"github.com/redis/go-redis/v9"
)

var client = redis.NewClient(config.RedisOptions)
var ctx = context.Background()

func Store(names []string, ttl time.Duration) error {
	key := utils.JSONTime(time.Now()).RoundHour().String()

	names = []string{}

	tmp := make([]interface{}, len(names))
	for i, v := range names {
		tmp[i] = v
	}
	var err error = nil
	if (len(names) == 0) {
		err = client.SAdd(ctx, key, []interface{}{0}...).Err()
	} else {
		err = client.SAdd(ctx, key, tmp...).Err()
	}
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
	result := make(map[int][]string)
	for startTime := time.Now().Add(time.Duration(-count + 1) * time.Hour); 
		startTime.Compare(time.Now()) <= 0; 
		startTime = startTime.Add(time.Hour) {

		key := utils.JSONTime(startTime).RoundHour().String()
		k := int(time.Time(utils.JSONTime(startTime).RoundHour()).UnixMilli())
		var err error
		result[k], err = client.SInter(ctx, key).Result()
		if err != nil {
			if err == redis.Nil {
				result[k] = []string{}
				continue
			}
			log.Default().Printf("Was not able to get Redis key '%d': %s. Skipping.", k, err)
			delete(result, k)
			continue
		}
	}

	return result, nil
}