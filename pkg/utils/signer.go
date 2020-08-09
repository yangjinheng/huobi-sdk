package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

// const
const (
	BasicDateFormat  = "2006-01-02T15:04:05"
	Algorithm        = "HmacSHA256"
	SignatureVersion = "2"
)

// Signer Signature Huobi meta
type Signer struct {
	Endpoint  string
	AccessKey string
	SecretKey string
}

func shouldEscape(c byte) bool {
	if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' || c == '_' || c == '-' || c == '~' || c == '.' {
		return false
	}
	return true
}

func escape(s string) string {
	hexCount := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if shouldEscape(c) {
			hexCount++
		}
	}
	if hexCount == 0 {
		return s
	}
	t := make([]byte, len(s)+2*hexCount)
	j := 0
	for i := 0; i < len(s); i++ {
		switch c := s[i]; {
		case shouldEscape(c):
			t[j] = '%'
			t[j+1] = "0123456789ABCDEF"[c>>4]
			t[j+2] = "0123456789ABCDEF"[c&15]
			j += 3
		default:
			t[j] = s[i]
			j++
		}
	}
	return string(t)
}

func hmacSha256base64(params, secret string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secret))
	_, err := mac.Write([]byte(params))
	if err != nil {
		return "", err
	}
	signByte := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signByte), nil
}

func sortQueryEncrypt(r *http.Request, accessKey string, secretKey string) error {
	var keys []string
	query := r.URL.Query()
	timestamp := time.Now().UTC().Format(BasicDateFormat)
	query.Add("AccessKeyId", accessKey)
	query.Add("SignatureMethod", Algorithm)
	query.Add("SignatureVersion", SignatureVersion)
	query.Add("Timestamp", timestamp)
	for key := range query {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var a []string
	for _, key := range keys {
		k := escape(key)
		sort.Strings(query[key])
		for _, v := range query[key] {
			kv := fmt.Sprintf("%s=%s", k, escape(v))
			a = append(a, kv)
		}
	}
	queryString := strings.Join(a, "&")
	stringToSign := fmt.Sprintf("%s\n%s\n%s\n%s", r.Method, r.URL.Host, r.URL.Path, queryString)
	signature, err := hmacSha256base64(stringToSign, secretKey)
	if err != nil {
		return err
	}
	signature = fmt.Sprintf("&Signature=%s", signature)
	queryString += signature
	r.URL.RawQuery = queryString
	return nil
}

// Sign Request Sign
func (s *Signer) Sign(r *http.Request) error {
	endpoint, err := url.Parse(s.Endpoint)
	if err != nil {
		return err
	}
	r.URL.Scheme, r.URL.Host = endpoint.Scheme, endpoint.Host
	sortQueryEncrypt(r, s.AccessKey, s.SecretKey)
	return nil
}
