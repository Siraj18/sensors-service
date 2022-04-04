package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/siraj18/sensor-checker/internal/ports"
	"github.com/sirupsen/logrus"
)

type handler struct {
	router    *chi.Mux
	logger    *logrus.Logger
	sensorSrv ports.SensorService
}

func NewHandler(sensorSrv ports.SensorService) *handler {
	return &handler{
		router:    chi.NewRouter(),
		logger:    logrus.New(),
		sensorSrv: sensorSrv,
	}
}

func (h *handler) getSensorsData(w http.ResponseWriter, r *http.Request) {
	data, err := h.sensorSrv.GetSensorsData()

	if err != nil {
		h.logger.Error(err.Error())
		return
	}

	json.NewEncoder(w).Encode(data)
}

func (handler *handler) InitRoutes() *chi.Mux {

	handler.router.Get("/getSensorsData", handler.getSensorsData)

	return handler.router
}
