package namaztime

import (
	"encoding/json"
	"fmt"
	"net/http"

	"marusya/internal/marusya"

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
	r.Post("/marusya", t.marusyaWebHook)
	r.Post("/alisa", t.alisaWebHook)
	return r
}

func (t *Transport) marusyaWebHook(w http.ResponseWriter, r *http.Request) {
	t.logger.Debug().Msg(fmt.Sprintf("request: %#v", r))

	var request *marusya.MarusyaRequest

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

	response := marusya.NewMarusyaResponse(msg, request)
	t.logger.Debug().Msg(fmt.Sprintf("response: %#v", r))

	render.JSON(w, r, response)
	return
}

func (t *Transport) alisaWebHook(w http.ResponseWriter, r *http.Request) {
	t.logger.Debug().Msg(fmt.Sprintf("request: %#v", r))
}
