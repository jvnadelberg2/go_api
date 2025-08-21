// Command identity-check connects to an Ory project using client-go,
// fetches the first available identity, and prints its ID together with
// either the username or email trait.
//
// Environment variables required:
//   ORY_PROJECT_URL  – Base URL of the Ory project (e.g. https://project.oryapis.com)
//   ORY_API_KEY      – API key with permission to list identities
//
// Output:
//   - If no identities exist, prints "0".
//   - Otherwise prints "<identity-id> <username-or-email>" to stdout.
//
// Intended use is as a quick diagnostic/utility tool to verify that
// the Ory API is reachable and contains at least one identity.

package main

import (
	"context"
	"fmt"
	"os"

	ory "github.com/ory/client-go"
)

// main builds a client from environment config, queries identities, and prints a summary.
func main() {
	// Build Ory client configuration from environment.
	cfg := ory.NewConfiguration()
	cfg.Servers = ory.ServerConfigurations{{URL: os.Getenv("ORY_PROJECT_URL")}}
	cfg.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("ORY_API_KEY"))

	c := ory.NewAPIClient(cfg)

	// Request the first page of identities, limit to 1.
	ids, _, err := c.IdentityAPI.ListIdentities(context.Background()).
		PerPage(1).
		Execute()
	if err != nil {
		panic(err)
	}

	// No identities found: print "0" and exit.
	if len(ids) == 0 {
		fmt.Println("0")
		return
	}

	// Extract username (preferred) or fall back to email from identity traits.
	uname := ""
	if m, ok := ids[0].Traits.(map[string]any); ok {
		if s, ok := m["username"].(string); ok && s != "" {
			uname = s
		} else if s, ok := m["email"].(string); ok {
			uname = s
		}
	}

	// Print identity ID and chosen trait.
	fmt.Printf("%s %s\n", ids[0].Id, uname)
}
