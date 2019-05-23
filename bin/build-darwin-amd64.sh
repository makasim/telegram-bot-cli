#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o build/telegram-bot-cli-darwin-amd64