package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const DefaultTimeout = 3.0

var (
	ErrPushoverFailed = errors.New("Pushover send failed")
)

func checkHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	url := os.Getenv("URL")
	if !strings.HasPrefix(url, "http") {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("URL is invalid. - %v", url),
			StatusCode: 200,
		}, nil
	}
	sec, err := strconv.ParseFloat(os.Getenv("TIMEOUT_SEC"), 32)
	if err != nil {
		sec = DefaultTimeout
	}
	timeout := time.Duration(sec) * time.Second
	log.Printf("Checking %s", url)

	if err := checkStatus(url, timeout); err != nil {
		err := pushOver(os.Getenv("PUSHOVER_APIKEY"),
			os.Getenv("PUSHOVER_USERKEY"), url+" is down!")
		if err != nil {
			log.Println("Pushover alert sent failed - " + err.Error())
			return events.APIGatewayProxyResponse{
				Body:       fmt.Sprintf("%s is down! Pushover alert sent also failed!", url),
				StatusCode: 500,
			}, ErrPushoverFailed
		}
		log.Println("Pushover alert sent successfully")
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("%s is down! Pushover alert sent successfully.", url),
			StatusCode: 200,
		}, nil
	}
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("%s is Good.", url),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(checkHandler)
}
