package namaztime

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rs/zerolog"
)

type Transport struct {
	service *Service
	logger  zerolog.Logger
}

func NewTransport(service *Service, logger zerolog.Logger) *Transport {
	return &Transport{
		service: service,
		logger:  logger,
	}
}

func (t *Transport) Handler() http.Handler {
	r := chi.NewRouter()
	r.Post("/webhook", t.webHookHandler)
	return r
}

func (t *Transport) webHookHandler(w http.ResponseWriter, r *http.Request) {
	var request *MarusyaRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	defer func() { _ = r.Body.Close() }()

	msg, err := t.service.GetNamazTimeMessage(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	render.JSON(w, r, NewMarusyaResponse(msg, request))
	return
}
