package util

import (
	"cache-service/model"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
)

func GetUUID() string {
	return uuid.New().String()
}

func DeleteAllItems(cache *cache.Cache) {
	cache.Flush()
}

func SetCache(cache *cache.Cache, data model.Data)  {
	cache.Set(GetUUID(), data.Value, -1)
}

//func GetDataFromCache(cache *cache.Cache) []model.Data {
//	var dataList []model.Data
//	items := cache.Items()
//	var c model.Data
//	for key, value := range items {
//		c = model.Data{
//			Key:   key,
//			Value: value.Object,
//		}
//		dataList = append(dataList, c)
//	}
//	return dataList
//}
