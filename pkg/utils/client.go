package utils

import (
	"net"
	"net/http"
	"net/url"
	"time"
)

// var
var (
	DefaultTransport *http.Transport
)

func init() {
	DefaultTransport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		IdleConnTimeout:        1 * time.Minute,
		TLSHandshakeTimeout:    10 * time.Second,
		ExpectContinueTimeout:  1 * time.Second,
		DisableKeepAlives:      false,
		MaxResponseHeaderBytes: 1 << 15,
	}
}

// ParseProxy "socks5://127.0.0.1:1080"
func ParseProxy(proxyURL string) (res func(*http.Request) (*url.URL, error), err error) {
	var purl *url.URL
	purl, err = url.Parse(proxyURL)
	if err != nil {
		return
	}
	res = http.ProxyURL(purl)
	return
}

// DefaultHTTPClient DefaultHTTPClient
func DefaultHTTPClient(proxyURL string) *http.Client {
	if proxyURL != "" {
		DefaultTransport.Proxy, _ = ParseProxy(proxyURL)
	}
	httpClient := &http.Client{
		Timeout:   10 * time.Second,
		Transport: DefaultTransport,
	}
	return httpClient
}
