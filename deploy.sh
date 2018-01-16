#!/bin/sh
GOOS=linux go build -o build/main
aws cloudformation package \
    --template-file template.yml \
    --s3-bucket exdit-lambda-deploy \
    --s3-prefix lambda-http-monitor \
    --output-template-file .template.yml
aws cloudformation deploy \
    --template-file .template.yml \
    --stack-name lambda-http-monitor \
    --capabilities CAPABILITY_IAM
