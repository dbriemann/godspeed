package godspeed

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Key uint32

type keyboard struct {
	keys [sdl.NUM_SCANCODES]State
	any  State
}

var (
	// Keyboard exposes all keyboard functionality.
	Keyboard = keyboard{}
)

func updateKeyboardStatus() {

}

func processKeyboardEvent(kbe *sdl.KeyboardEvent) {
	fmt.Println("scancode", kbe.Keysym.Scancode)
	fmt.Println("keycode", kbe.Keysym.Sym)
	fmt.Println("mod", kbe.Keysym.Mod)
	fmt.Println("repeat", kbe.Repeat)

}
