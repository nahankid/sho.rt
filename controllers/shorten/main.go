package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"sho.rt/lib"
)

// ShortenRequest struct
type ShortenRequest struct {
	URL        string `json:"url"`
	ExpiryDays int    `json:"expiry_days"`
}

// URL struct
type URL struct {
	APIKey    string `json:"-"`
	ShortURL  string `json:"short_url"`
	LongURL   string `json:"long_url"`
	ExpiresAt string `json:"-"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req ShortenRequest
	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return lib.APIResponse(http.StatusUnprocessableEntity, err.Error())
	}
	if req.URL == "" {
		return lib.APIResponse(http.StatusBadRequest, "Url is missing. Nothing to shorten")
	}

	surl, err := lib.RandomHex(3)
	if err != nil {
		return lib.APIResponse(http.StatusInternalServerError, err.Error())
	}

	expiry := 7
	if req.ExpiryDays > 0 {
		expiry = req.ExpiryDays
	}

	url := URL{
		APIKey:    request.RequestContext.Identity.APIKey,
		LongURL:   req.URL,
		ShortURL:  surl,
		ExpiresAt: strconv.FormatInt(time.Now().AddDate(0, 0, expiry).Unix(), 10),
	}

	log.Println("url is", url)

	err = putItem(&url)
	if err != nil {
		return lib.APIResponse(http.StatusInternalServerError, err.Error())
	}

	res, err := json.Marshal(url)
	if err != nil {
		return lib.APIResponse(http.StatusInternalServerError, err.Error())
	}

	return lib.APIResponse(http.StatusCreated, string(res))
}

func main() {
	lambda.Start(handler)
}
