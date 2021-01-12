package util

import (
	"bytes"
	"cache-service/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UtilsTestSuite struct {
	suite.Suite
	context  *gin.Context
	recorder *httptest.ResponseRecorder
}

func TestCacheHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UtilsTestSuite))
}

func (suite *UtilsTestSuite) SetupTest() {
	suite.recorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.recorder)
}

func (suite *UtilsTestSuite) TearDownTest() {
	suite.recorder.Flush()
}

func (suite *UtilsTestSuite) TestDecodeRequestBody() {
	body := `{"data":{"name":"test name"}}`

	req := ioutil.NopCloser(bytes.NewReader([]byte(body)))
	request := http.Request{Body: req}

	expected := map[string]interface{}{"data": map[string]interface{}{"name": "test name"}}
	requestBody, err := DecodeRequestBody(&request)

	suite.Equal(expected, requestBody.Value)
	suite.Nil(err)
	suite.NotNil(requestBody)
}

func (suite *UtilsTestSuite) TestDecodeRequestBodyShouldReturnError() {
	body := ``

	req := ioutil.NopCloser(bytes.NewReader([]byte(body)))
	request := http.Request{Body: req}

	expected := model.Data{}
	requestBody, err := DecodeRequestBody(&request)

	suite.Equal(expected, requestBody)
	suite.NotNil(err)
}

func (suite *UtilsTestSuite) TestShouldGetPaginationFromRequest() {
	actualPagination, err := GetPagination("1", "10")
	expectedPagination := model.Pagination{
		PageNumber: 1,
		PageSize:   10,
	}

	suite.Equal(expectedPagination, actualPagination)
	suite.NotNil(actualPagination)
	suite.Nil(err)
}

func (suite *UtilsTestSuite) TestGetPaginationFromRequestReturnErrorForInvalidPageNumber() {
	actualPagination, err := GetPagination("one", "10")
	expected := model.Pagination{}

	suite.Equal(expected, actualPagination)
	suite.NotNil(err)
}

func (suite *UtilsTestSuite) TestGetPaginationFromRequestReturnErrorForInvalidPageSize() {
	actualPagination, err := GetPagination("10", "one")
	expected := model.Pagination{}

	suite.Equal(expected, actualPagination)
	suite.NotNil(err)
}
