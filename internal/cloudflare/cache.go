package cloudflare

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/fzdy-zz/proxypool/config"
)

func test() {
	api, err := cloudflare.New(config.Config.CFKey, config.Config.CFKey)
	if err != nil {
		log.Fatal(err)
	}

	// Fetch the zone ID
	id, err := api.ZoneIDByName(config.Config.Domain)
	if err != nil {
		log.Fatal(err)
	}

	// Fetch zone details
	zone, err := api.ZoneDetails(context.Background(), id)
	if err != nil {
		log.Fatal(err)
	}
	// Print zone details
	fmt.Println(zone)

}
