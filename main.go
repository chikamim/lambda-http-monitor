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

const DefaultTimeout = 2.5

var (
	ErrPushoverSent = errors.New("Pushover sent failed")
)

func checkHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	url := os.Getenv("URL")
	if !strings.HasPrefix(url, "http") {
		return response(fmt.Sprintf("URL is invalid. - %v", url), 200), nil
	}
	sec, err := strconv.ParseFloat(os.Getenv("TIMEOUT_SEC"), 32)
	if err != nil {
		sec = DefaultTimeout
	}
	timeout := time.Duration(sec) * time.Second
	log.Printf("Checking %v\n", url)

	err = checkStatus(url, timeout)
	if err != nil {
		err = pushOver(os.Getenv("PUSHOVER_APIKEY"), os.Getenv("PUSHOVER_USERKEY"), url+" is down!")
		if err != nil {
			log.Println("Pushover alert sent failed - " + err.Error())
			return response(fmt.Sprintf("%s is down! Pushover alert sent also failed!", url), 500), ErrPushoverSent
		}
		log.Println("Pushover alert sent successfully")
		return response(fmt.Sprintf("%s is down! Pushover alert sent successfully.", url), 200), nil
	}
	log.Printf("%s is OK.\n", url)
	return response(fmt.Sprintf("%s is OK.", url), 200), nil
}

func response(message string, status int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       message,
		StatusCode: status,
	}
}

func main() {
	lambda.Start(checkHandler)
}
