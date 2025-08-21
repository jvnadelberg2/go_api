package main

import (
	"context"
	"fmt"
	"os"

	ory "github.com/ory/client-go"
)

func main() {
	cfg := ory.NewConfiguration()
	cfg.Servers = ory.ServerConfigurations{{URL: os.Getenv("ORY_PROJECT_URL")}}
	cfg.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("ORY_API_KEY"))

	c := ory.NewAPIClient(cfg)
	ids, _, err := c.IdentityAPI.ListIdentities(context.Background()).PerPage(1).Execute()
	if err != nil {
		panic(err)
	}
	if len(ids) == 0 {
		fmt.Println("0")
		return
	}

	uname := ""
	if m, ok := ids[0].Traits.(map[string]any); ok {
		if s, ok := m["username"].(string); ok && s != "" {
			uname = s
		} else if s, ok := m["email"].(string); ok {
			uname = s
		}
	}
	fmt.Printf("%s %s\n", ids[0].Id, uname)
}
