#!/usr/bin/env bash

#env GOOS=linux GOARCH=amd64 go build -o ottosfa-api-web

#scp -i ~/.ssh/LightsailDefaultKey-ap-southeast-1-new.pem -P 22 ottosfa-api-web ubuntu@13.228.25.85:/home/ubuntu

env GOOS=linux GOARCH=amd64 go build -o ottosfa-api-web

scp -i ~/.ssh/devsfa.priv -P 22 ottosfa-api-web devsfa@34.101.141.240:/home/devsfa

#env GOOS=linux GOARCH=amd64 go build -o ottosfa-api-web
#scp -i ~/.ssh/devofin.priv -P 22 ottosfa-api-web devofin@34.101.141.240:~/
#ssh devofin@34.101.141.240 -i ~/.ssh/devofin.priv