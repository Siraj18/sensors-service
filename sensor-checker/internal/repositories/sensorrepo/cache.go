package sensorrepo

import (
	"errors"

	"github.com/siraj18/sensor-checker/internal/models"
	"github.com/siraj18/sensor-checker/pkg/cachedb"
)

type CacheRepository struct {
	cache *cachedb.CacheDb
}

func NewCacheRepository(cache *cachedb.CacheDb) *CacheRepository {
	return &CacheRepository{
		cache: cache,
	}
}

func (rep *CacheRepository) Get(key string) (*models.SensorsData, error) {

	item := rep.cache.Get(key)

	if item == nil {
		return nil, cachedb.ErrorItemNotFound
	}

	sensorData, ok := item.(*models.SensorsData)
	if !ok {
		return nil, errors.New("error when get item")
	}

	return sensorData, nil
}

func (rep *CacheRepository) Set(key string, value *models.SensorsData) {
	rep.cache.Set(key, value)
}
