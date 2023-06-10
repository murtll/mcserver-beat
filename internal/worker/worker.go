package worker

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/murtll/mcserver-beat/config"
	"github.com/murtll/mcserver-beat/internal/service"
	"github.com/murtll/mcserver-beat/internal/entities"
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
		var body entities.MinetoolsPollingResponse
		err = json.Unmarshal(rawBody, &body)
		if err != nil {
			log.Default().Printf("Was not able to unmarshal %s: %s.", rawBody, err)
			continue
		}
		log.Default().Printf("Got %v players. Saving.", body.Playerlist)
		err = service.Store(body.Playerlist, config.EntryTTL)
		if err != nil {
			log.Default().Printf("Was not able to store data: %s", err)
		}
	}
}