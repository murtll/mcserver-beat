package service

import (
	"time"

	"github.com/murtll/mcserver-beat/internal/entities"
	"github.com/murtll/mcserver-beat/internal/utils"
	"github.com/murtll/mcserver-beat/internal/repository"
)

func Store(count int, ttl time.Duration) error {
	return repository.Store(count, ttl)
}

func Load(count int) (*entities.GraphResponse, error) {
	rawData, err := repository.Load(count)
	if err != nil {
		return nil, err
	}

	max := 0
	data := make([]entities.PlayerCount, 0)
	for k, v := range rawData {
		if max < v {
			max = v
		}
		data = append(data, entities.PlayerCount{Time: utils.JSONTime(time.UnixMilli(int64(k))), Number: v})
	}
	return &entities.GraphResponse{
		Max: max,
		Data: data,
	}, nil
}