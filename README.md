# ory-go-demo

Minimal Go demo using the [Ory Cloud Admin API](https://www.ory.sh/docs/reference/api) to list identities.  
This example shows how to authenticate with an Ory project using an API key and read back the first identity.

---

## Purpose
- Demonstrate calling the Ory Admin API from Go.
- Show use of environment variables for project URL and API key.
- Print the ID and username/email of the first available identity.

---

## Requirements
- Go 1.22+
- An Ory Cloud project (URL + API key)
- See shared prerequisites in `docs/PREREQS.md`

---

## Quickstart

Clone and enter the repo:

    git clone <your-repo-url> ory-go-demo
    cd ory-go-demo

Set environment variables:

    export ORY_PROJECT_URL="https://<slug>.projects.oryapis.com"
    export ORY_API_KEY="ory_pat_xxx..."

Run the demo:

    ./run.sh

---

## Expected output

On success, the program prints the identity UUID and either `username` or `email`:

    8f268dd1-3ec3-43ed-83f3-cfd469fb8906 jonnadelberg

If no identities exist:

    0

---

## Environment

- `ORY_PROJECT_URL` – Ory project base URL (e.g. `https://<slug>.projects.oryapis.com`)
- `ORY_API_KEY` – personal API key for the project

---

## Endpoints used
- `GET /admin/identities?per_page=1` – list identities (first page)

---

## Error handling
- Exits non-zero on error.
- Typical issues:
  - 401 — invalid or missing API key
  - 404 — incorrect project URL
  - Network/DNS problems

---

## Notes for extension
- Print all identities, not just the first.
- Display all traits (e.g. email, profile).
- Create identities with `POST /admin/identities`.
- Integrate into a larger Go service.

---

## License
MIT
