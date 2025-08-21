# Security: ory-go-demo

## Scope
This demo uses the Ory Admin API with an API key. It is not a production system.  
The purpose is to demonstrate SDK usage, not to provide a hardened implementation.

## Authentication
- Requires environment variables:
  - `ORY_PROJECT_URL` — Ory Cloud project base URL
  - `ORY_API_KEY` — personal API key created in Ory Cloud dashboard
- The API key is passed as a Bearer token in the `Authorization` header.

## Key handling
- API keys grant administrative rights. Treat them as secrets.
- Never commit real keys into version control.
- Use environment variables or a secrets manager when running outside a demo.
- Rotate keys regularly in Ory Cloud and scope them as narrowly as possible.

## Data sensitivity
- Calls to `ListIdentities` return identity IDs and traits (username/email).
- Do not expose this output beyond the demo context.
- For production, enforce least-privilege access and audit API key usage.

## Threat model (demo)
- Anyone with the API key can query and manage identities in the project.
- Loss of the key compromises the project.
- Demo code does not include access controls, rate limits, or logging.

## Responsible disclosure
This repository is a demo. Do not report vulnerabilities in this demo project.  
For Ory platform security issues, see [https://www.ory.sh/security](https://www.ory.sh/security).
