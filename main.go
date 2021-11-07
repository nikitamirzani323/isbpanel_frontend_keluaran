package main

import (
	"log"

	"github.com/nikitamirzani323/isbpanel_frontend_keluaran/router"
)

func main() {
	app := router.Init()
	log.Fatal(app.Listen(":7072"))
}
