package main

import (
	"context"
	"fmt"
	"os"

	ory "github.com/ory/client-go"
)

func main() {
	base := os.Getenv("ORY_PROJECT_URL")
	key := os.Getenv("ORY_API_KEY")

	cfg := ory.NewConfiguration()
	cfg.Servers = ory.ServerConfigurations{{URL: base}}
	cfg.AddDefaultHeader("Authorization", "Bearer "+key)

	c := ory.NewAPIClient(cfg)
	ids, _, err := c.IdentityAPI.ListIdentities(context.Background()).PerPage(1).Execute()
	if err != nil {
		panic(err)
	}
	if len(ids) == 0 {
		fmt.Println("0")
		return
	}

	id := ids[0].Id
	uname := ""
	if m, ok := ids[0].Traits.(map[string]any); ok {
		if v, ok := m["username"].(string); ok {
			uname = v
		} else if v, ok := m["email"].(string); ok {
			uname = v
		}
	}
	fmt.Printf("%s %s\n", id, uname)
}
