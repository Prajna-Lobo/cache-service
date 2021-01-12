package service

import (
	"cache-service/model"
	"cache-service/repository"
	"cache-service/util"
	"errors"
	"github.com/patrickmn/go-cache"
)

type cacheService struct {
	cache      *cache.Cache
	repository repository.StorageRepository
}

type CacheService interface {
	StoreData(data model.Data) error
	FetchFromCache(pagination model.Pagination) (model.Response, error)
}

func NewCacheService(cache *cache.Cache, repository repository.StorageRepository) cacheService {
	return cacheService{
		cache:      cache,
		repository: repository,
	}
}

func (service cacheService) StoreData(data model.Data) error {
	if len(data.Key) != 0 {
		util.SetCache(service.cache, data)
		err := service.repository.InsertData(data)
		if err != nil {
			return err
		}
	} else {
		return errors.New("no data found to store")
	}
	return nil
}

func (service cacheService) FetchFromCache(pagination model.Pagination) (model.Response, error) {
	res, err := service.repository.ReloadDataFromDb(service.cache, pagination)
	if err != nil {
		return model.Response{}, err
	}
	return res, nil
}
