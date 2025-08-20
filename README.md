# ory-go-demo

Minimal Go demo using net/http that validates an Ory session via the Frontend API and protects a route. If the session is missing or inactive, the app redirects to the Ory login UI.

## What it does
- Reads a browser session cookie issued by Ory
- Calls Ory Frontend API to validate the session
- Serves `/ping` only when the session is active; otherwise redirects to `${ORY_BASE_URL}/ui/login`

## Requirements
- Go 1.22+
- See shared prerequisites in `docs/PREREQS.md`

## Quickstart
git clone <your-repo-url> ory-go-demo
cd ory-go-demo
go mod tidy
go run .

## Environment
export ORY_BASE_URL=https://playground.projects.oryapis.com
export ORY_SESSION_COOKIE=ory_session_playground

## Obtain a session
1. Open `$ORY_BASE_URL/ui/login` and sign in or register.
2. Copy the cookie named by `$ORY_SESSION_COOKIE`.

## Test
curl 'http://localhost:8080/ping' -b "$ORY_SESSION_COOKIE=<cookie_value>"

Expected:
{"message":"pong"}

If not authenticated, the server responds with a redirect to `$ORY_BASE_URL/ui/login`.

## VS Code
Use a launch config that sets `ORY_BASE_URL` and `ORY_SESSION_COOKIE`.

## GitHub
- Build: `go build ./...`
- Smoke: `go run .` and curl `/ping` with a valid cookie in CI

## Project structure
- `main.go`
- `go.mod`
- `docs/PREREQS.md`
- `docs/ory-go-demo-architecture.svg`

## Diagram
See `docs/ory-go-demo-architecture.svg`.

## License
MIT
