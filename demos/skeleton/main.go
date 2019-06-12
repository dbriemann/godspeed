package main

import (
	"github.com/dbriemann/godspeed"
)

type game struct {
	state float64
}

func (g *game) init() {
}

func (g *game) update(dt float64) {
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
