package main

import (
	"command_line/app"
	"log"
	"os"
)

func main() {
	app := app.Gerar()
	if error := app.Run(os.Args); error != nil {
		log.Fatal(error)
	}

}
