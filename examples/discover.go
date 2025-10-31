package main

import (
	"context"
	"fmt"
	"log"
	
	"github.com/atlas402/index/core"
)

func main() {
	ctx := context.Background()
	index := core.New("https://facilitator.payai.network")
	
	options := &core.DiscoveryOptions{
		Category: "AI",
		Network:  "base",
	}
	
	services, err := index.Discover(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Found %d services\n", len(services))
}


