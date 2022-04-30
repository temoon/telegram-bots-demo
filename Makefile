#!make

include .env
export $(shell sed 's/=.*//' .env)

SHELL=/bin/bash -o pipefail

.PHONY: all bucket webhook updates

all: bucket

bucket:
	rm -f demo-bot.zip
	zip -r demo-bot.zip router vars go.mod go.sum yandex.go

webhook:
	curl --request POST --url https://api.telegram.org/bot${BOT_TOKEN}/setWebhook --header 'content-type: application/json' --data '{"url": "${BOT_URL}"}'

updates:
	curl --request POST --url https://api.telegram.org/bot${BOT_TOKEN}/deleteWebhook --header 'content-type: application/json' --data '{}'
