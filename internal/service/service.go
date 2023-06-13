package service

import (
	"sort"
	"time"

	"github.com/murtll/mcserver-beat/internal/entities"
	"github.com/murtll/mcserver-beat/internal/repository"
	"github.com/murtll/mcserver-beat/internal/utils"
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
	data := make([]entities.PlayerCount, count)
	i := 0
	for k, names := range rawData {
		nameCount := len(names)
		if max < nameCount {
			max = nameCount
		}
		data[i] = entities.PlayerCount{
			Time: utils.JSONTime(time.UnixMilli(int64(k))), 
			Number: nameCount,
			Players: names,
		}
		i++
	}
	sort.Slice(data, func(i, j int) bool {
		return time.Time(data[i].Time).Compare(time.Time(data[j].Time)) == -1
	})

	return &entities.GraphResponse{
		Max: max,
		Data: data,
	}, nil
}