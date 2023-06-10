package service

import (
	"time"

	"github.com/murtll/mcserver-beat/internal/entities"
	"github.com/murtll/mcserver-beat/internal/utils"
	"github.com/murtll/mcserver-beat/internal/repository"
)

func Store(names []string, ttl time.Duration) error {
	return repository.Store(names, ttl)
}

func Load(count int) (*entities.GraphResponse, error) {
	rawData, err := repository.Load(count)
	if err != nil {
		return nil, err
	}

	max := 0
	data := make([]entities.PlayerCount, 0)
	for k, names := range rawData {
		nameCount := len(names)
		if max < nameCount {
			max = nameCount
		}
		data = append(data, 
			entities.PlayerCount{
				Time: utils.JSONTime(time.UnixMilli(int64(k))), 
				Number: nameCount,
				Players: names,
			})
	}
	return &entities.GraphResponse{
		Max: max,
		Data: data,
	}, nil
}