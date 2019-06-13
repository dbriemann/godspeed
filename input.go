package godspeed

import (
	"github.com/veandco/go-sdl2/sdl"
)

// processEvents receives events from the sdl event queue and
// maps them to the internal representation. It updates all
// corresponding engine variables.
func processEvents() {
	// Update internal status depending on the last frame.
	Mouse.updateStatus()
	Keyboard.updateStatus()

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
