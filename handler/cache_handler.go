package handler

import (
	"cache-service/constants"
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
// @Success 200 {} string  "fetch data"
// @Failure 400 {} string  "Bad request when the given format in invalid"
// @Failure 500 {} string "All service errors"
// @Router /api/cache-service/v1/data [post]
func (handler cacheHandler) StoreCacheData(c *gin.Context) {
	cache, err := util.DecodeRequestBody(c.Request)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	serviceErr := handler.cacheService.StoreData(cache)
	if serviceErr != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.AbortWithStatus(http.StatusOK)
}

// Fetch data from Cache service godoc
// @Tags Cache service
// @Summary Fetch data from cache using pagination
// @Param page_num query int true "Page number"
// @Param page_size query int true "Page size"
// @Success 200 {object} model.Response "fetches the data successfully"
// @Failure 400 {} string "Bad request page_num or page_size invalid"
// @Failure 500 {} string "all service errors"
// @Router /api/cache-service/v1/data [get]
func (handler cacheHandler) FetchCacheData(c *gin.Context) {
	pageNum := c.Query(constants.PageNumber)
	pageSize := c.Query(constants.PageNumber)
	pagination, err := util.GetPagination(pageNum, pageSize)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	results, err := handler.cacheService.FetchFromCache(pagination)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.AbortWithStatusJSON(http.StatusOK, results)
}
