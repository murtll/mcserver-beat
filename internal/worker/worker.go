package worker

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/murtll/mcserver-beat/config"
	"github.com/murtll/mcserver-beat/internal/service"
)

func Start() {
	for {
		time.Sleep(config.PollingInterval)
		res, err := http.Get(config.PollingURL)
		if err != nil {
			log.Default().Printf("Was not able to contact %s: %s.", config.PollingURL, err)
			continue
		}
		rawBody := make([]byte, res.ContentLength)
		res.Body.Read(rawBody)
		log.Default().Printf("Got data: %s.", string(rawBody))
		var body map[string]any
		err = json.Unmarshal(rawBody, &body)
		if err != nil {
			log.Default().Printf("Was not able to unmarshal %s: %s.", rawBody, err)
			continue
		}
		data := int(body[config.PollingSchema].(float64))
		log.Default().Printf("Got '%d' count. Saving.", data)
		err = service.Store(data, config.EntryTTL)
		if err != nil {
			log.Default().Printf("Was not able to store data: %s", err)
		}
	}
}