package main

import (
	"databuseasyway/internal/cli"
	"log"
)

func main() {
	c := cli.NewClient()
	if err := c.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
