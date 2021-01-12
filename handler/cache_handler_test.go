package handler

import (
	"bytes"
	"cache-service/mocks"
	"cache-service/model"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CacheHandlerTestSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
	context  *gin.Context
	recorder *httptest.ResponseRecorder
	service  *mocks.MockCacheService
	handler  CacheHandler
}

func TestCacheHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CacheHandlerTestSuite))
}

func (suite *CacheHandlerTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.recorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.recorder)
	suite.service = mocks.NewMockCacheService(suite.mockCtrl)
	suite.handler = NewCacheHandler(suite.service)
}

func (suite *CacheHandlerTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
	suite.recorder.Flush()
}

func (suite *CacheHandlerTestSuite) TestStoreCacheDataShouldStoreDataInCacheAndDb() {
	data := `{"data":{"name":"testName"}}`
	suite.context.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString(data))

	suite.service.EXPECT().StoreData(gomock.Any()).Return(nil)
	suite.handler.StoreCacheData(suite.context)

	_, _ = ioutil.ReadAll(suite.recorder.Body)
	suite.Equal(http.StatusOK, suite.recorder.Code)
}

func (suite *CacheHandlerTestSuite) TestStoreCacheDataShouldThrowBadRequest() {
	data := ``
	suite.context.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString(data))

	suite.service.EXPECT().StoreData(gomock.Any()).Return(nil)
	suite.handler.StoreCacheData(suite.context)

	_, _ = ioutil.ReadAll(suite.recorder.Body)
	suite.Equal(http.StatusBadRequest, suite.recorder.Code)
}

func (suite *CacheHandlerTestSuite) TestStoreCacheDataShouldThrowError() {
	data := `{"data":{"name":"testName"}}`
	suite.context.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString(data))

	suite.service.EXPECT().StoreData(gomock.Any()).Return(errors.New("error while storing data"))
	suite.handler.StoreCacheData(suite.context)

	_, _ = ioutil.ReadAll(suite.recorder.Body)
	suite.Equal(http.StatusInternalServerError, suite.recorder.Code)
}

func (suite *CacheHandlerTestSuite) TestFetchDataFromCacheShouldReturnSuccess() {
	suite.context.Request, _ = http.NewRequest("GET", "?page_num=1&page_size=10", nil)
	data := model.Response{
		Data: []model.Data{
			{
				"abc",
				"data",
			},
			{
				"test2",
				"data2",
			},
		},
		Pagination: model.PaginationResponse{
			TotalPageCount: 2,
		},
	}
	suite.service.EXPECT().FetchFromCache(gomock.Any()).Return(data, nil)
	suite.handler.FetchCacheData(suite.context)

	_, _ = ioutil.ReadAll(suite.recorder.Body)

	suite.Equal(http.StatusOK, suite.recorder.Code)
}

func (suite *CacheHandlerTestSuite) TestFetchDataFromCacheShouldReturnBadRequest() {
	suite.context.Request, _ = http.NewRequest("GET", "?page_num=one", nil)
	data := model.Response{
		Data: []model.Data{
			{
				"abc",
				"data",
			},
			{
				"test2",
				"data2",
			},
		},
		Pagination: model.PaginationResponse{
			TotalPageCount: 2,
		},
	}
	suite.service.EXPECT().FetchFromCache(gomock.Any()).Return(data, nil)
	suite.handler.FetchCacheData(suite.context)

	_, _ = ioutil.ReadAll(suite.recorder.Body)

	suite.Equal(http.StatusBadRequest, suite.recorder.Code)
}

func (suite *CacheHandlerTestSuite) TestFetchDataFromCacheShouldReturnInternalServerError() {
	suite.context.Request, _ = http.NewRequest("GET", "?page_num=1&page_size=10", nil)
	suite.service.EXPECT().FetchFromCache(gomock.Any()).Return(model.Response{}, errors.New("something went wrong"))
	suite.handler.FetchCacheData(suite.context)
	_, _ = ioutil.ReadAll(suite.recorder.Body)

	suite.Equal(http.StatusInternalServerError, suite.recorder.Code)
}
