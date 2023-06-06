package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/murtll/mcserver-beat/config"
	"github.com/murtll/mcserver-beat/internal/entities"
	"github.com/murtll/mcserver-beat/internal/service"
)

var mux = http.NewServeMux()

func handleGraphInfo(w http.ResponseWriter, r *http.Request) {
	log.Default().Printf("%s %s%s?%s", r.Method, r.Host, r.URL.Path, r.URL.RawQuery)
	w.Header().Add("content-type", "application/json")
	if r.Method != http.MethodGet {
		data, err := json.Marshal(entities.RestError{
			Message: "Method not allowed.",
		})
		if err != nil {
			log.Default().Printf(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(data)
		return
	}

	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil {
		count = config.EntryNumber
	}

	graphInfo, err := service.Load(count)
	if err != nil {
		log.Default().Printf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(graphInfo)
	if err != nil {
		log.Default().Printf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func handleHealthInfo(w http.ResponseWriter, r *http.Request) {
	log.Default().Printf("%s %s%s?%s", r.Method, r.Host, r.URL.Path, r.URL.RawQuery)
	if r.Method != http.MethodGet {
		data, err := json.Marshal(entities.RestError{
			Message: "Method not allowed.",
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(data)
		return
	}
	w.Write([]byte("Ok."))
}

var server = &http.Server{
	Addr: config.ListenAddr,
	Handler: mux,
}

func init() {
	mux.HandleFunc(config.ListenPath, handleGraphInfo)
	mux.HandleFunc("/_healthz", handleHealthInfo)
}

func Start() error {
	log.Default().Printf("Starting server on http://%s", server.Addr)
	return server.ListenAndServe()
}