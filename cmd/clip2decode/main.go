package main

import (
	"fmt"
	"github/tr3mor/clip2decode/internal/app/clip2decode"
)

func main() {
	c := clip2decode.NewClipboard()
	app := clip2decode.NewApplication(c)
	fmt.Println(app.DecodeData())
}
