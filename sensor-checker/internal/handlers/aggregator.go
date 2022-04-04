package handlers

import (
	"context"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/siraj18/sensor-checker/internal/models"
	"github.com/siraj18/sensor-checker/internal/ports"
	"github.com/sirupsen/logrus"
)

type aggregator struct {
	urls      []string
	client    *http.Client
	logger    *logrus.Logger
	sensorSrv ports.SensorService
}

func NewAggregator(urls []string, sensorSrv ports.SensorService, timeout time.Duration) *aggregator {
	return &aggregator{
		urls: urls,
		client: &http.Client{
			Timeout: timeout,
		},
		sensorSrv: sensorSrv,
		logger:    logrus.New(),
	}
}

// Идет начальный сбор данных
// Затем запускается цикл с определеннным таймаутом
func (ag *aggregator) InitCheck(ctx context.Context, duration time.Duration) {
	ag.aggregate()

	ticker := time.NewTicker(duration)

	go func() {
		for {
			select {
			case <-ticker.C:
				ag.aggregate()
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()

}

func (ag *aggregator) aggregate() {

	if len(ag.urls) == 0 {
		return
	}

	results := make(chan int, len(ag.urls))
	errs := make(chan error, len(ag.urls))

	var wg sync.WaitGroup
	wg.Add(len(ag.urls))

	for _, url := range ag.urls {
		go func(url string) {
			defer wg.Done()

			answer, err := ag.checkUrl(url)
			if err != nil {
				errs <- err
				return
			}

			results <- answer
		}(url)
	}

	wg.Wait()
	close(results)
	close(errs)

	count := len(results)
	var sum int

	for result := range results {
		sum = sum + result
	}

	var finalAnswer models.SensorsData
	finalAnswer.DataIsFull = true

	if len(errs) > 0 {
		finalAnswer.DataIsFull = false
		for err := range errs {
			ag.logger.Error("service error:", err.Error())
		}
	}

	// Если все сервисы легли
	if count == 0 {
		ag.sensorSrv.AddSensorsData(&finalAnswer)
		return
	}

	arithMean := float64(sum) / float64(count)

	finalAnswer.Value = arithMean

	ag.sensorSrv.AddSensorsData(&finalAnswer)

}

func (ag *aggregator) checkUrl(url string) (int, error) {
	resp, err := ag.client.Get(url)

	if err != nil {
		return 0, err
	}

	content, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return 0, err
	}

	answer, _ := strconv.Atoi(string(content))

	resp.Body.Close()

	return answer, nil
}
