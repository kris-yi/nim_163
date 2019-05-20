package server

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"github.com/google/go-querystring/query"
	"math/big"
	"net/url"
	"strconv"
	"time"
)

type Request interface {
	GetQuery() (url.Values, error)
	GetMethod() string
	GetUrl(baseUrl string) string
	GetHeader(appKey, appSecret string) map[string]string
}

type NimRequest struct {
	Api       string
	QueryData interface{}
}

func (request *NimRequest) GetQuery() (url.Values, error) {
	return query.Values(request.QueryData)
}

func (request *NimRequest) GetMethod() string {
	return "POST"
}

func (request *NimRequest) GetUrl(baseUrl string) string {
	return baseUrl + request.Api
}

func (request *NimRequest) GetHeader(appKey, appSecret string) map[string]string {
	nonce := createNonce(32)
	curTime := strconv.FormatInt(time.Now().Unix(), 10)

	return map[string]string{
		"Content-Type": "application/x-www-form-urlencoded;charset=utf-8",
		"AppKey":       appKey,
		"Nonce":        nonce,
		"CurTime":      curTime,
		"CheckSum": func(appSecret, nonce, curTime string) string {
			str := appSecret + nonce + curTime
			h := sha1.New()
			h.Write([]byte(str))
			hash := h.Sum(nil)
			return fmt.Sprintf("%x", hash)
		}(appSecret, nonce, curTime),
	}
}

func createNonce(length int) string {
	str := "12345567890abcdefghijklmnopqrstuvwxyz"
	nonce := ""
	if length < 1 {
		return ""
	}
	strLen := int64(len(str))
	for i := 0; i < length; i++ {
		randStr, _ := rand.Int(rand.Reader, big.NewInt(strLen))
		start := randStr.Int64()
		end := start + 1
		nonce += str[start:end]
	}
	return nonce
}
