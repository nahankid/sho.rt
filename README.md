# Sho.rt

![Go](https://github.com/nahankid/shortie/workflows/Go/badge.svg)
![GitHub](https://img.shields.io/github/license/nahankid/sho.rt)
[![Maintainability](https://api.codeclimate.com/v1/badges/eb06a36a6fcda7abc6f2/maintainability)](https://codeclimate.com/github/nahankid/sho.rt/maintainability)
![GitHub repo size](https://img.shields.io/github/repo-size/nahankid/sho.rt)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/nahankid/sho.rt)


Sho.rt is a Fast URL shortener, written in Go. This app uses compute only for shortening a URL. It doesn't use any compute for fetching URLs, redirecting to long urls and for cleanup of expired URLs. All business logic is handled at the Amazon API Gateway level. 

Entire appliation stack is 1-click deploy on AWS and uses: 

- AWS Route 53
- AWS Certificate Manager
- AWS API Gateway
- AWS Lambda
- DynamoDB
- AWS Serverless Application Model (SAM)

## Quick start

## **Shorten a long url**

```POST https://c1ix.me``` 

### Parameters

| Name         | Type     | Description                                              |
| ------------ | ---------| -------------------------------------------------------- | 
| url          | string   | Long URL to be shortened.                                | 
| expiry_days  | int      | Optional. Days for the short link to expire. Default: 7  | 


### Response

| Name         | Type     | Description                                              |
| ------------ | ---------| -------------------------------------------------------- | 
| url          | string   | Short URL                                                | 


## Contributing to Shortie

Fork, fix, then send me a pull request.

## License

MIT
