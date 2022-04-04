package main

import (
	"context"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/siraj18/sensor-checker/internal/handlers"
	"github.com/siraj18/sensor-checker/internal/repositories/sensorrepo"
	"github.com/siraj18/sensor-checker/internal/server"
	"github.com/siraj18/sensor-checker/internal/services/sensorsrv"
	"github.com/siraj18/sensor-checker/pkg/cachedb"
	"github.com/sirupsen/logrus"
)

func main() {
	urlString := os.Getenv("services_urls")
	httpClientTimeouts, _ := strconv.Atoi(os.Getenv("http_client_timeouts_in_seconds"))
	aggregateTimeouts, _ := strconv.Atoi(os.Getenv("aggregate_repeat_timeouts_in_seconds"))
	address := os.Getenv("address")

	urls := strings.Split(urlString, ";") // Разделяем url сервисов по символу ;

	cache := cachedb.NewCache()
	cacheRepository := sensorrepo.NewCacheRepository(cache)

	sensorService := sensorsrv.NewSensorsService(cacheRepository)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	aggregator := handlers.NewAggregator(urls, sensorService, time.Second*time.Duration(httpClientTimeouts))
	aggregator.InitCheck(ctx, time.Second*time.Duration(aggregateTimeouts))

	handler := handlers.NewHandler(sensorService)
	srv := server.NewServer(address, handler.InitRoutes(), time.Second*10)

	if err := srv.Run(); err != nil {
		logrus.Fatal(err.Error())
	}
}
