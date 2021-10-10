package main

import (
	"log"

	"github.com/nikitamirzani323/gofiber_backendtogel/router"
)

func main() {
	app := router.Init()
	log.Fatal(app.Listen(":7071"))
}
