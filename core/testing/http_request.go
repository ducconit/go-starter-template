package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	assert "github.com/stretchr/testify/require"
)

type Request struct {
	t       *testing.T
	headers map[string]string
	body    io.Reader
	query   url.Values
}

func NewRequestBuilder(t *testing.T) *Request {
	return &Request{
		t:       t,
		headers: map[string]string{},
		query:   url.Values{},
	}
}

func (r *Request) WithHeader(key string, value string) *Request {
	r.headers[key] = value
	return r
}

func (r *Request) WithContentType(value string) *Request {
	return r.WithHeader("Content-Type", value)
}

func (r *Request) WithAuthorization(value string) *Request {
	return r.WithHeader("Authorization ", value)
}

func (r *Request) WithBearerAuthorization(value string) *Request {
	return r.WithAuthorization("Bearer " + value)
}

func (r *Request) WithBasicAuthorization(value string) *Request {
	return r.WithAuthorization("Basic " + value)
}

func (r *Request) WithBody(body io.Reader) *Request {
	r.body = body
	return r
}

func (r *Request) WithBodyString(body string) *Request {
	r.body = bytes.NewReader([]byte(body))
	return r
}

func (r *Request) WithBodyJson(body interface{}) *Request {
	if data, err := json.Marshal(body); err == nil {
		r.body = bytes.NewReader(data)
	}
	return r.WithContentType("application/json")
}

func (r *Request) WithQuery(key string, value string) *Request {
	r.query.Add(key, value)
	return r
}

func (r *Request) Request(method string, url string) *http.Request {
	req, err := http.NewRequest(method, fmt.Sprintf("%s?%s", url, r.query.Encode()), r.body)
	assert.NoError(r.t, err)
	for k, val := range r.headers {
		req.Header.Add(k, val)
	}
	return req
}

func (r *Request) Get(url string) *http.Request {
	return r.Request(http.MethodGet, url)
}

func (r *Request) Post(url string) *http.Request {
	return r.Request(http.MethodPost, url)
}

func (r *Request) Put(url string) *http.Request {
	return r.Request(http.MethodPut, url)
}

func (r *Request) Patch(url string) *http.Request {
	return r.Request(http.MethodPatch, url)
}

func (r *Request) Delete(url string) *http.Request {
	return r.Request(http.MethodDelete, url)
}
