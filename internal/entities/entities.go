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

type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

type PlayerCount struct {
	Number int `json:"number"`
	Time   utils.JSONTime `json:"time"`
	Players []string `json:"players"`
}

type MinetoolsPollingResponse struct {
	MaxPlayers int
	Motd string
	Playerlist []string
	Players int
	Plugins []string
	Software string
	Version string
	Status string `json:"status"`
}