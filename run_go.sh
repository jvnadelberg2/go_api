#!/usr/bin/env bash
set -e
cd "$(dirname "$0")"

export ORY_PROJECT_URL="https://dreamy-grothendieck-hgdo1j9d82.projects.oryapis.com"
export ORY_API_KEY="ory_pat_ClPM7rh56HERpspLFFHNyDtdJCsCpXru"

if [ ! -f go.mod ]; then
  go mod init go_api
fi

go get github.com/ory/client-go@v1.21.5
go mod tidy
go run main.go
