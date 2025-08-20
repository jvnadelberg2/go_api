# Prerequisites

## Ory environment
- Base URL: use the Ory Network playground or your project.
  - Playground: https://playground.projects.oryapis.com
  - Your project: https://<PROJECT_SLUG>.projects.oryapis.com
- For session-based flows (Kratos), create a user via `/ui/login` or `/ui/registration`.
- For permission checks (Keto), ensure you have a Personal Access Token.

## Environment variables
- ORY_BASE_URL (default: https://playground.projects.oryapis.com)
- ORY_SESSION_COOKIE (playground uses ory_session_playground)
- ORY_TOKEN (required only for Python demo)

## VS Code
- Enable Go or Python extension as appropriate.
- Use a per-repo .vscode/launch.json to set env vars.

## GitHub Actions
- Cache dependencies.
- Add a smoke test job that runs the app/script with ORY_BASE_URL set.
- Store secrets in Actions secrets.

## Troubleshooting
- 401 or redirect loops: invalid session or cookie name.
- 404 on Ory endpoints: check ORY_BASE_URL.
- TLS/proxy: avoid intercepting proxies.
