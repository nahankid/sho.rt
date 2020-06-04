package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-south-1"))

// Add a url record to DynamoDB.
func putItem(url *URL) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("ShortenerAPI-LinkTable"),
		Item: map[string]*dynamodb.AttributeValue{
			"api_key": {
				S: aws.String(url.APIKey),
			},
			"short_url": {
				S: aws.String(url.ShortURL),
			},
			"long_url": {
				S: aws.String(url.LongURL),
			},
			"expires_at": {
				N: aws.String(url.ExpiresAt),
			},
		},
	}

	_, err := db.PutItem(input)
	return err
}
