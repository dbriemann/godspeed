package godspeed

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// Private engine variables and functions.
var (
	version       = "0.0.1-alpha"
	isInitialized = false
	tickTime      = 10 * time.Millisecond
	frameTime     = tickTime.Seconds()
	running       = false

	window *sdl.Window
	fps    uint
)

var (
	// Core engine functions which should be re-implemented by
	// the user.

	// Init is run once before the game is started.
	Init = func() {

	}

	// Update is called for every logic step of the engine.
	// The parameter dt contains the time in seconds since
	// the last call to Update.
	Update = func(dt float64) {

	}

	// Draw is called for every draw step of the engine.
	// It will use the active RenderTarget to draw onto.
	Draw = func() {

	}

	// Running returns true if the engine is running and false
	// otherwise. The default function checks for sdl quit event.
	Running = func() bool {
		// Poll events and swap buffers, etc.
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return false
			}
		}
		return true
	}
)

var (
	// Secondary engine functions that can be safely ignored
	// by the user most of the time.

	// Loop runs the actual gameloop of a godspeed game.
	Loop = func() {
		var lag float64

		start := time.Now()
		last := start

		up := start
		var frames uint

		for Running() {
			current := time.Now()
			elapsed := current.Sub(last)
			last = current
			lag += elapsed.Seconds()

			dirty := false
			// Call update and draw functions.
			for lag >= frameTime {
				Update(frameTime)
				lag -= frameTime
				dirty = true
			}

			// If an Update happened call Draw.
			// Otherwise skip draw step.
			if dirty {
				Draw()
				frames++
			}

			snap := current.Sub(up).Seconds()

			if snap > 1.0 {
				fps = uint(math.Round(float64(frames) / snap))
				up = current
				frames = 0
			}

			sleepFor := (frameTime - current.Sub(last).Seconds()) * 0.99
			time.Sleep(time.Duration(sleepFor * float64(time.Second)))
		}
	}
)

// Run takes care of engine initialization and starts the
// actual gameloop function Loop.
func Run() {
	// Run provided game Init function.
	Init()

	// Init SDL.
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic("failed to init SDL: " + err.Error())
	}
	defer sdl.Quit()

	// Create SDL window.
	window, err := sdl.CreateWindow(
		Config.Window.Title,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		int32(Config.Window.Width),
		int32(Config.Window.Height),
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		panic("failed to create SDL window: " + err.Error())
	}
	defer window.Destroy()

	running = true
	Loop()
}
