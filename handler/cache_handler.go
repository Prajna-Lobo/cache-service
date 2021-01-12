package handler

import (
	"cache-service/constants"
	"cache-service/model"
	"cache-service/service"
	"cache-service/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type cacheHandler struct {
	cacheService service.CacheService
}

type CacheHandler interface {
	StoreCacheData(ctx *gin.Context)
	FetchCacheData(ctx *gin.Context)
}

func NewCacheHandler(service service.CacheService) cacheHandler {
	return cacheHandler{
		cacheService: service,
	}
}

// Store data in Cache service godoc
// @Tags Cache service
// @Summary Store data in cache with persistence backup in DB
// @Accept json
// @Param Data body interface{} true "Any valid json Data can be provided"
// @Success 201 {} string  "store data"
// @Failure 400 {object} model.Error  "ErrorCode: ERR_BAD_REQUEST"
// @Failure 500 {object} model.Error "ErrorCode: ERR_INTERNAL_SERVER"
// @Router /api/cache-service/v1/data [post]
func (handler cacheHandler) StoreCacheData(c *gin.Context) {
	cache, err := util.DecodeRequestBody(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.GetBadRequestError())
	}
	serviceErr := handler.cacheService.StoreData(cache)
	if serviceErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.GetInternalServerError())
	}
	c.AbortWithStatus(http.StatusCreated)
}

// Fetch data from Cache service godoc
// @Tags Cache service
// @Summary Fetch data from cache using pagination
// @Param page_num query int true "Page number"
// @Param page_size query int true "Page size"
// @Success 200 {object} model.Response "fetches the data from cache"
// @Failure 400 {object} model.Error "ErrorCode: ERR_BAD_REQUEST"
// @Failure 500 {object} model.Error "ErrorCode: ERR_INTERNAL_SERVER"
// @Router /api/cache-service/v1/data [get]
func (handler cacheHandler) FetchCacheData(c *gin.Context) {
	pageNum := c.Query(constants.PageNumber)
	pageSize := c.Query(constants.PageNumber)
	pagination, err := util.GetPagination(pageNum, pageSize)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.GetBadRequestError())
	}
	results, err := handler.cacheService.FetchFromCache(pagination)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.GetInternalServerError())
	}
	c.AbortWithStatusJSON(http.StatusOK, results)
}
