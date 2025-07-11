package testing

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

type RestAssured struct {
	t        *testing.T
	resp     *http.Response
	respBody string
}

func NewRestAssured(t *testing.T, resp *http.Response) *RestAssured {
	respBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	return &RestAssured{
		t:        t,
		resp:     resp,
		respBody: string(respBody),
	}
}

func (r *RestAssured) Status(httpStatusCode int) *RestAssured {
	assert.Equal(r.t, httpStatusCode, r.resp.StatusCode)
	return r
}

func (r *RestAssured) Header(key string, expected interface{}) *RestAssured {
	assert.EqualValues(r.t, expected, r.resp.Header.Get(key))
	return r
}

func (r *RestAssured) HasSize(key string, expected int) *RestAssured {
	v := gjson.Get(r.respBody, key)
	assert.EqualValues(r.t, expected, len(v.Array()), fmt.Sprintf("Expected size of %v is %v, but got: %v", key, expected, len(v.Array())))
	return r
}

func (r *RestAssured) Body(key string, expected interface{}) *RestAssured {
	v := gjson.Get(r.respBody, key)
	assert.EqualValues(r.t, expected, v.Value(), fmt.Sprintf("Expected value of key %v is %v, but got: %v", key, expected, v.Value()))
	return r
}

func (r *RestAssured) BodyFn(key string, expectedFn func(value interface{})) *RestAssured {
	v := gjson.Get(r.respBody, key)
	expectedFn(v.Value())
	return r
}

func (r *RestAssured) String() string {
	return r.respBody
}

func (r *RestAssured) Extract(key string, result interface{}) {
	assert.NotNil(r.t, result)
	v := gjson.Get(r.respBody, key)
	err := json.Unmarshal([]byte(v.String()), result)
	assert.NoError(r.t, err)
}
