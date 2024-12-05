package main

import (
	"context"
	"log"

	"github.com/stevekroll/example/cmd"
)

func main() {
	ctx := context.Background()
	err := cmd.ExecuteContext(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
