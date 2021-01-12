package repository

import (
	"cache-service/model"
	"cache-service/util"
	"fmt"
	"github.com/patrickmn/go-cache"
	"gopkg.in/mgo.v2"
)

type storageRepository struct {
	collection *mgo.Collection
}

type StorageRepository interface {
	InsertData(data model.Data) error
	ReloadDataFromDb(cache *cache.Cache, pagination model.Pagination) (model.Response,error)
}

func NewStorageRepository(collection *mgo.Collection) storageRepository {
	return storageRepository{
		collection: collection,
	}
}

func (repository storageRepository) InsertData(data model.Data) error {
	err := repository.collection.Insert(data)
	return err
}

func (repository storageRepository) ReloadCacheFromDb(cache *cache.Cache) error {
	var data []model.Data
	util.DeleteAllItems(cache)
	err := repository.collection.Find(nil).All(&data)
	if err != nil {
		return err
	}
	for _, result := range data {
		fmt.Printf("%+v", result)
		util.SetCache(cache, result)
	}
	return nil
}

func (repository storageRepository) ReloadDataFromDb(cache *cache.Cache, pagination model.Pagination) (model.Response, error) {
	results, err := repository.getDataFromDb(pagination)
	if err != nil {
		return model.Response{}, err
	}
	util.DeleteAllItems(cache)
	for _, result := range results.Data {
		util.SetCache(cache, result)

	}
	return results, nil
}

func (repository storageRepository) getDataFromDb(pagination model.Pagination) (model.Response, error) {
	pageResponse, err := repository.getPageCount(pagination)
	if err != nil {
		return model.Response{}, err
	}
	data, err := repository.queryToGetData(pagination)
	if err != nil {
		return model.Response{}, err
	}
	response := model.Response{Data: data, Pagination: pageResponse}
	return response, nil
}

func (repository storageRepository) getPageCount(pagination model.Pagination) (model.PaginationResponse, error) {
	collectionCount, err := repository.collection.Count()
	if err != nil {
		return model.PaginationResponse{}, err
	}
	return model.PaginationResponse{
		TotalPageCount: collectionCount / pagination.PageSize,
	}, nil
}

func (repository storageRepository) queryToGetData(pagination model.Pagination) ([]model.Data, error) {
	var data []model.Data
	query := repository.collection.Find(nil)
	query = query.Skip((pagination.PageNumber - 1) * pagination.PageSize).Limit(pagination.PageSize)
	queryErr := query.All(&data)
	return data, queryErr
}
