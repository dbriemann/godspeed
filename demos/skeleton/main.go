package main

import (
	"fmt"

	"github.com/dbriemann/godspeed"
)

type game struct {
	state float64
}

func (g *game) init() {
}

func (g *game) update(dt float64) {
	if godspeed.Keyboard.KeyDown(4) {
		fmt.Println("A DOWN")
	} else if godspeed.Keyboard.KeyPressed(4) {
		fmt.Println("A PRESSED")
	} else if godspeed.Keyboard.KeyReleased(4) {
		fmt.Println("A RELEASED")
	}
}

func (g *game) draw() {
}

func main() {
	g := game{}
	godspeed.Init = g.init
	godspeed.Update = g.update
	godspeed.Draw = g.draw

	godspeed.Run()
}
