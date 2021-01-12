package util

import (
	"cache-service/model"
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
)

func DecodeRequestBody(r *http.Request) (model.Data, error) {
	var cache model.Data
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var data interface{}
	err := decoder.Decode(&data)

	if reflect.ValueOf(data).IsValid() {
		cache = model.Data{
			Key:   GetUUID(),
			Value: data,
		}
	}
	return cache, err
}

func GetPagination(pNum, pSize string) (model.Pagination, error) {
	pageNum, err := strconv.Atoi(pNum)
	if err != nil {
		return model.Pagination{}, err
	}
	pageSize, err := strconv.Atoi(pSize)
	if err != nil {
		return model.Pagination{}, err
	}

	return model.Pagination{
		PageNumber: pageNum,
		PageSize:   pageSize,
	}, nil
}
