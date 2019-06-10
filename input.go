package godspeed

import (
	"github.com/veandco/go-sdl2/sdl"
)

// State describes the state of a Button or a Key.
type State uint8

const (
	// Up describes that a key or button is up.
	Up State = iota
	// Down describes that a key or button is down.
	Down
	// Pressed describes that a key or button was just pressed.
	Pressed
	// Released descibes that a key or button was just released.
	Released
)

var (
	stateToStr = map[State]string{
		Up:       "up",
		Down:     "down",
		Pressed:  "pressed",
		Released: "released",
	}
)

// StateStr returns a human-readable description for a State.
func StateStr(s State) string {
	return stateToStr[s]
}

// processEvents receives events from the sdl event queue and
// maps them to the internal representation. It updates all
// corresponding engine variables.
func processEvents() {
	// Update internal status depending on the last frame.
	updateMouseStatus()

	// Poll events from sdl event queue.
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.MouseButtonEvent:
			mbe := event.(*sdl.MouseButtonEvent)
			processMouseButtonEvent(mbe)
		case *sdl.QuitEvent:
			running = false
		}
	}
}
