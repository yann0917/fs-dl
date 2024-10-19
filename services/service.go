package services

import (
	"github.com/go-resty/resty/v2"
)

var (
	baseURL   = "https://gateway-api.dushu365.com"
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"
)

type Service struct {
	client *resty.Client
}

func NewService() *Service {
	headers := map[string]string{
		"Accept":          "application/json, text/plain, */*",
		"Content-Type":    "application/json;charset=UTF-8",
		"User-Agent":      userAgent,
		"Priority":        "u=1, i",
		"Reqentryption":   "AES",
		"X-Dushu-App-Plt": "3",
		"X-Dushu-App-Ver": "1.0.0",
	}
	client := resty.New()
	//client.SetDebug(true)
	client.SetBaseURL(baseURL).
		SetHeaders(headers)

	return &Service{client: client}
}
