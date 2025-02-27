package main

import (
	"log"
	"github.com/henrytw/bitfinex-api-go/v2/rest"
	"github.com/davecgh/go-spew/spew"
)


func main() {
	apikey := ""
	secKey := ""

	c := rest.NewClient().Credentials(apikey, secKey)

	positions, err := c.Positions.All()

	if err != nil {
		log.Fatalf("getting wallet %s", err)
	}

	spew.Dump(positions)
}

