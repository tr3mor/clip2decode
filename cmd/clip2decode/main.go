package main

import (
	"github/tr3mor/clip2decode/internal/app/clip2decode"
	"log"
)

func main() {
	c := clip2decode.NewClipboard()
	app := clip2decode.NewApplication(c)
	output, err := app.DecodeData()
	if err != nil {
		log.Fatalf("Execution failed due to: %s", err)
	}
	println(output)
}
