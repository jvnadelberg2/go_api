# Architecture: ory-go-demo

![Ory Go Demo Architecture](./ory-go-demo-architecture.svg)

_Figure: The Ory Go Demo architecture shows a minimal Go client calling the Ory Cloud Admin API to read identity data.  
On the left, the Go program (`main.go`) runs locally with `ORY_PROJECT_URL` and `ORY_API_KEY` set in the environment.  
The client makes an HTTPS request to list identities (`GET /admin/identities?per_page=1`).  
On the right, the Ory Cloud project responds with JSON containing identity objects, including `id` and `traits` (such as `username` or `email`).  

The client prints the first identity’s UUID and a human-friendly identifier (username preferred, else email). Identity storage and management remain in Ory’s managed services; the client holds no long-term state beyond the API key provided via environment variables._

---

## Purpose
This is a minimal demo that shows how a Go program can connect to an Ory Cloud project using an API key and read back identity information.

- Authenticate with Ory Cloud via API key  
- Retrieve identity data  
- Print the first identity’s UUID and username/email

## Components
- **Client**: `main.go` (Go 1.22+, uses `github.com/ory/client-go`)  
- **Ory Project**: `https://<project-slug>.projects.oryapis.com`  
- **Network**: HTTPS from client → Ory  

## Environment parameters
- `ORY_PROJECT_URL` — Ory project base URL (e.g., `https://<slug>.projects.oryapis.com`)  
- `ORY_API_KEY` — personal access token with permissions for the Admin API  

## Ory endpoints used
- `GET /admin/identities?per_page=1` — list identities (fetch one for the demo)  

## Execution steps
1. Read configuration  
   Load `ORY_PROJECT_URL` and `ORY_API_KEY` from the environment.  
2. Initialize client  
   Create `ory.NewAPIClient` using `ORY_PROJECT_URL`; set `Authorization: Bearer <ORY_API_KEY>`.  
3. List identities  
   Call `IdentityAPI.ListIdentities(...).PerPage(1).Execute()`.  
4. Output  
   If an identity exists, print: `<identity_id> <username_or_email>`; otherwise print `0`.  

## Inputs / outputs
- Inputs: `ORY_PROJECT_URL`, `ORY_API_KEY`  
- Outputs:  
  - `<uuid> <username-or-email>`  
  - or `0` if no identities exist  

## Error handling
- Non-2xx responses cause the program to exit with an error.  
- Common issues:  
  - 401 — invalid or missing API key  
  - 404 — incorrect project URL  
  - Network errors — DNS, proxy, firewall  

## Security note
- Do not hard-code API keys in source; supply them via environment variables.  
- Rotate keys regularly and restrict scope according to least privilege.  

## Minimal run
go mod tidy  
export ORY_PROJECT_URL="https://<slug>.projects.oryapis.com"  
export ORY_API_KEY="ory_pat_xxx..."  
go run main.go  

## Notes for extension
- Iterate through all identities (pagination) and print additional traits.  
- Create identities with `POST /admin/identities`; update/delete for lifecycle management.  
- Add structured error handling and timeouts.  
- Integrate into a service (e.g., Gin/Echo/Fiber) with HTTP endpoints.  

## Summary
- Go client authenticates with Ory Cloud using an API key.  
- Calls the Admin API to list identities.  
- Prints a concise identity summary for verification.  
