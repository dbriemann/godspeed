package godspeed

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Key uint32

type keyboard struct {
	keys [sdl.NUM_SCANCODES]State
	any  bool
}

func (k *keyboard) KeyState(key Key) State {
	return k.keys[key]
}

// TODO get key descriptor

var (
	// Keyboard exposes all keyboard functionality.
	Keyboard = keyboard{}
)

func updateKeyboardStatus() {
	Keyboard.any = false
	for i := 0; i < len(Keyboard.keys); i++ {
		switch Keyboard.keys[i] {
		case Pressed:
			// If key was pressed last frame set it to down.
			Keyboard.keys[i] = Down
		case Released:
			// If key was release last frame set it to up.
			Keyboard.keys[i] = Up
		}
	}
}

func processKeyboardEvent(kbe *sdl.KeyboardEvent) {
	switch kbe.State {
	case sdl.PRESSED:
		Keyboard.any = true
		Keyboard.keys[kbe.Keysym.Scancode] = Pressed
	case sdl.RELEASED:
		Keyboard.keys[kbe.Keysym.Scancode] = Released
	}

	// TODO mod
	// fmt.Println("mod", kbe.Keysym.Mod)
}
