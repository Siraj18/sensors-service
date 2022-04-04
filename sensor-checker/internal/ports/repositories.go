package ports

import "github.com/siraj18/sensor-checker/internal/models"

type CacheRepository interface {
	Get(key string) (*models.SensorsData, error)
	Set(key string, value *models.SensorsData)
}
