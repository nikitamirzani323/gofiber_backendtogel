package main

import (
	"log"

	"bitbucket.org/isbtotogroup/frontendagen_svelte/router"
)

func main() {
	app := router.Init()
	log.Fatal(app.Listen(":7071"))
}
