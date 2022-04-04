package sensorsrv

import (
	"github.com/siraj18/sensor-checker/internal/models"
	"github.com/siraj18/sensor-checker/internal/ports"
)

type service struct {
	sensorsRepo ports.CacheRepository
}

func NewSensorsService(sensorsRepo ports.CacheRepository) *service {
	return &service{
		sensorsRepo: sensorsRepo,
	}
}

func (s *service) GetSensorsData() (*models.SensorsData, error) {
	data, err := s.sensorsRepo.Get(models.KEY)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *service) AddSensorsData(sensorData *models.SensorsData) {
	s.sensorsRepo.Set(models.KEY, sensorData)
}
