package testing

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

type MockGinContext struct {
	ctx *gin.Context
}

func NewMockGinContext(w *httptest.ResponseRecorder) *MockGinContext {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	return &MockGinContext{
		ctx: ctx,
	}
}

func (m *MockGinContext) Context() *gin.Context {
	return m.ctx
}

func (m *MockGinContext) WithContext(key string, value interface{}) *MockGinContext {
	m.ctx.Set(key, value)
	return m
}

func (m *MockGinContext) WithParams(params gin.Params) *MockGinContext {
	m.ctx.Params = params
	return m
}

func (m *MockGinContext) WithQuery(key string, value string) *MockGinContext {
	query := m.ctx.Request.URL.Query()
	query.Add(key, value)
	m.ctx.Request.URL.RawQuery = query.Encode()
	return m
}

func (m *MockGinContext) WithRequest(request *http.Request) *MockGinContext {
	m.ctx.Request = request
	return m
}
