package router

import (
	"cache-service/configuration"
	_ "cache-service/docs"
	"cache-service/factory"
	"cache-service/handler"
	"cache-service/repository"
	"cache-service/service"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
)

func SetUpRouter(config configuration.ConfigData) *gin.Engine {
	router := gin.Default()

	cacheClient, err := factory.GetCache()
	if err != nil {
		log.Fatal("unable to create a cache", err)
	}
	mongoCollection := factory.GetMongoCollection(&config)
	storageRepository := repository.NewStorageRepository(mongoCollection)
	cacheService := service.NewCacheService(cacheClient, storageRepository)
	cacheHandler := handler.NewCacheHandler(cacheService)

	routerGroup := router.Group("/api/cache-service/v1/")
	{
		routerGroup.GET("/health", HealthCheck)
		routerGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		cacheGroup := routerGroup.Group("/data")
		{
			cacheGroup.POST("", cacheHandler.StoreCacheData)
			cacheGroup.GET("", cacheHandler.FetchCacheData)
		}
	}
	return router
}

func HealthCheck(c *gin.Context) {
	c.AbortWithStatusJSON(200, "Application is up and running")
}
