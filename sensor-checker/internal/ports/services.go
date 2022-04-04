package ports

import "github.com/siraj18/sensor-checker/internal/models"

type SensorService interface {
	GetSensorsData() (*models.SensorsData, error)
	AddSensorsData(sensorData *models.SensorsData)
}
