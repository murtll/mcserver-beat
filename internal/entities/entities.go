package entities

import (
	"github.com/murtll/mcserver-beat/internal/utils"
)

type RestError struct {
	Message string `json:"error"`
}

type GraphResponse struct {
	Max  int `json:"max"`
	Data []PlayerCount `json:"data"`
}

type PlayerCount struct {
	Number int `json:"number"`
	Time   utils.JSONTime `json:"time"`
}
