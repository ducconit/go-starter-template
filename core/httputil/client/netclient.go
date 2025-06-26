package client

import (
	"crypto/tls"
	"crypto/x509"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

type HTTPClientSettings struct {
	Connect          time.Duration
	ConnKeepAlive    time.Duration
	ExpectContinue   time.Duration
	IdleConn         time.Duration
	MaxAllIdleConns  int
	MaxHostIdleConns int
	ResponseHeader   time.Duration
	TLSHandshake     time.Duration
	CustomCerts      []*x509.Certificate // for custom server CAs
	ClientCerts      []tls.Certificate   // for mTLS
}

// NewHTTP create http client with htt2Enable configuration
func NewHTTP(httpSettings HTTPClientSettings, http2Enable bool) (*http.Client, error) {
	var client http.Client
	tr := &http.Transport{
		ResponseHeaderTimeout: httpSettings.ResponseHeader,
		Proxy:                 http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			KeepAlive: httpSettings.ConnKeepAlive,
			Timeout:   httpSettings.Connect,
		}).DialContext,
		MaxIdleConns:          httpSettings.MaxAllIdleConns,
		IdleConnTimeout:       httpSettings.IdleConn,
		TLSHandshakeTimeout:   httpSettings.TLSHandshake,
		MaxIdleConnsPerHost:   httpSettings.MaxHostIdleConns,
		ExpectContinueTimeout: httpSettings.ExpectContinue,
	}

	var rootCAPool *x509.CertPool
	if len(httpSettings.CustomCerts) > 0 {
		// create a Certificate pool to hold one or more CA certificates
		rootCAPool = x509.NewCertPool()
		for _, cert := range httpSettings.CustomCerts {
			rootCAPool.AddCert(cert)
		}
	}

	tr.TLSClientConfig = &tls.Config{
		RootCAs:      rootCAPool,
		Certificates: httpSettings.ClientCerts,
	}

	// So client makes HTTP/2 requests
	if http2Enable {
		err := http2.ConfigureTransport(tr)
		if err != nil {
			return &client, err
		}
	}

	return &http.Client{
		Transport: tr,
	}, nil
}
