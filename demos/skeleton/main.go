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
	s := godspeed.Mouse.ButtonState(godspeed.ButtonX2)
	fmt.Println(s)
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
