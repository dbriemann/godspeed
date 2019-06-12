package godspeed

import (
	"github.com/veandco/go-sdl2/sdl"
)

// State describes the state of a Button or a Key.
type State uint8

func (s State) String() string {
	switch s {
	case Up:
		return "up"
	case Released:
		return "released"
	case Down:
		return "down"
	case Pressed:
		return "pressed"
	default:
		return "unknown"
	}
}

const (
	// Up describes that a key or button is up.
	Up State = iota
	// Released descibes that a key or button was just released.
	Released
	// Down describes that a key or button is down.
	Down
	// Pressed describes that a key or button was just pressed.
	Pressed
)

// processEvents receives events from the sdl event queue and
// maps them to the internal representation. It updates all
// corresponding engine variables.
func processEvents() {
	// Update internal status depending on the last frame.
	updateMouseStatus()
	updateKeyboardStatus()

	// Poll events from sdl event queue.
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.MouseButtonEvent:
			mbe := event.(*sdl.MouseButtonEvent)
			processMouseButtonEvent(mbe)
		case *sdl.MouseWheelEvent:
			mwe := event.(*sdl.MouseWheelEvent)
			processMouseWheelEvent(mwe)
		case *sdl.MouseMotionEvent:
			mme := event.(*sdl.MouseMotionEvent)
			processMouseMotionEvent(mme)
		case *sdl.KeyboardEvent:
			kbe := event.(*sdl.KeyboardEvent)
			processKeyboardEvent(kbe)
		case *sdl.QuitEvent:
			running = false
		}
	}
}
