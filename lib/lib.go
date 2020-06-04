package lib

import (
	"crypto/rand"
	"encoding/hex"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

// RandomHex returns a n*2 digit random hex
func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// APIResponse returns APIGatewayProxyResponse
func APIResponse(code int, msg string) (events.APIGatewayProxyResponse, error) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json; charset=UTF-8"
	headers["Access-Control-Allow-Origin"] = "*"

	log.Printf("%d - %s\n", code, msg)

	return events.APIGatewayProxyResponse{
		Body:       msg,
		StatusCode: code,
		Headers:    headers,
	}, nil
}
