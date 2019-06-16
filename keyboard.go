package godspeed

import (
	"fmt"

	"github.com/kr/pretty"
	"github.com/veandco/go-sdl2/sdl"
)

// Key is the internal type to represent a key.
type Key uint32

func (key Key) String() string {
	return sdl.GetKeyName(sdl.GetKeyFromScancode(internal2Scancode[key]))
}

const (
	KeyUnknown Key = iota
	Key0
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9
	KeyA
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyF13
	KeyF14
	KeyF15
	KeyF16
	KeyF17
	KeyF18
	KeyF19
	KeyF20
	KeyF21
	KeyF22
	KeyF23
	KeyF24
	KeyRight
	KeyLeft
	KeyDown
	KeyUp
	KeySpace
	KeyEscape
	KeyEnter
	KeyTab
	KeyBackspace
	KeyInsert
	KeyDelete
	KeyPageUp
	KeyPageDown
	KeyPrintScreen
	KeyPause
	KeyApostrophe
	KeyComma
	KeyMinus
	KeyPeriod
	KeySlash
	KeySemicolon
	KeyEqual
	KeyLeftBracket
	KeyBackslash
	KeyRightBracket
	KeyGraveAccent
	KeyWWW
	KeyHome
	KeyEnd
	KeyCapsLock
	KeyScrollLock
	KeyNumLock
	KeyNum0
	KeyNum1
	KeyNum2
	KeyNum3
	KeyNum4
	KeyNum5
	KeyNum6
	KeyNum7
	KeyNum8
	KeyNum9
	KeyNumDecimal
	KeyNumDivide
	KeyNumMultiply
	KeyNumSubtract
	KeyNumAdd
	KeyNumEnter
	KeyNumEqual
	KeyLeftShift
	KeyLeftControl
	KeyLeftAlt
	KeyLeftSuper
	KeyRightShift
	KeyRightControl
	KeyRightAlt
	KeyRightSuper
	KeyMenu
	KeyLast
)

var (
	scancode2Internal [sdl.NUM_SCANCODES]Key
	internal2Scancode [KeyLast]sdl.Scancode
)

type keyboard struct {
	lastInput [KeyLast]bool
	input     [KeyLast]bool
	any       bool
}

func (kb *keyboard) init() {
	// Create a mapping from internal codes to sdl scancodes for all
	// defined keys.
	internal2Scancode[KeyUnknown] = sdl.GetScancodeFromKey(sdl.K_UNKNOWN)
	internal2Scancode[Key0] = sdl.GetScancodeFromKey(sdl.K_0)
	internal2Scancode[Key1] = sdl.GetScancodeFromKey(sdl.K_1)
	internal2Scancode[Key2] = sdl.GetScancodeFromKey(sdl.K_2)
	internal2Scancode[Key3] = sdl.GetScancodeFromKey(sdl.K_3)
	internal2Scancode[Key4] = sdl.GetScancodeFromKey(sdl.K_4)
	internal2Scancode[Key5] = sdl.GetScancodeFromKey(sdl.K_5)
	internal2Scancode[Key6] = sdl.GetScancodeFromKey(sdl.K_6)
	internal2Scancode[Key7] = sdl.GetScancodeFromKey(sdl.K_7)
	internal2Scancode[Key8] = sdl.GetScancodeFromKey(sdl.K_8)
	internal2Scancode[Key9] = sdl.GetScancodeFromKey(sdl.K_9)
	internal2Scancode[KeyA] = sdl.GetScancodeFromKey(sdl.K_a)
	internal2Scancode[KeyB] = sdl.GetScancodeFromKey(sdl.K_b)
	internal2Scancode[KeyC] = sdl.GetScancodeFromKey(sdl.K_c)
	internal2Scancode[KeyD] = sdl.GetScancodeFromKey(sdl.K_d)
	internal2Scancode[KeyE] = sdl.GetScancodeFromKey(sdl.K_e)
	internal2Scancode[KeyF] = sdl.GetScancodeFromKey(sdl.K_f)
	internal2Scancode[KeyG] = sdl.GetScancodeFromKey(sdl.K_g)
	internal2Scancode[KeyH] = sdl.GetScancodeFromKey(sdl.K_h)
	internal2Scancode[KeyI] = sdl.GetScancodeFromKey(sdl.K_i)
	internal2Scancode[KeyJ] = sdl.GetScancodeFromKey(sdl.K_j)
	internal2Scancode[KeyK] = sdl.GetScancodeFromKey(sdl.K_k)
	internal2Scancode[KeyL] = sdl.GetScancodeFromKey(sdl.K_l)
	internal2Scancode[KeyM] = sdl.GetScancodeFromKey(sdl.K_m)
	internal2Scancode[KeyN] = sdl.GetScancodeFromKey(sdl.K_n)
	internal2Scancode[KeyO] = sdl.GetScancodeFromKey(sdl.K_o)
	internal2Scancode[KeyP] = sdl.GetScancodeFromKey(sdl.K_p)
	internal2Scancode[KeyQ] = sdl.GetScancodeFromKey(sdl.K_q)
	internal2Scancode[KeyR] = sdl.GetScancodeFromKey(sdl.K_r)
	internal2Scancode[KeyS] = sdl.GetScancodeFromKey(sdl.K_s)
	internal2Scancode[KeyT] = sdl.GetScancodeFromKey(sdl.K_t)
	internal2Scancode[KeyU] = sdl.GetScancodeFromKey(sdl.K_u)
	internal2Scancode[KeyV] = sdl.GetScancodeFromKey(sdl.K_v)
	internal2Scancode[KeyW] = sdl.GetScancodeFromKey(sdl.K_w)
	internal2Scancode[KeyX] = sdl.GetScancodeFromKey(sdl.K_x)
	internal2Scancode[KeyY] = sdl.GetScancodeFromKey(sdl.K_y)
	internal2Scancode[KeyZ] = sdl.GetScancodeFromKey(sdl.K_z)
	internal2Scancode[KeyF1] = sdl.GetScancodeFromKey(sdl.K_F1)
	internal2Scancode[KeyF2] = sdl.GetScancodeFromKey(sdl.K_F2)
	internal2Scancode[KeyF3] = sdl.GetScancodeFromKey(sdl.K_F3)
	internal2Scancode[KeyF4] = sdl.GetScancodeFromKey(sdl.K_F4)
	internal2Scancode[KeyF5] = sdl.GetScancodeFromKey(sdl.K_F5)
	internal2Scancode[KeyF6] = sdl.GetScancodeFromKey(sdl.K_F6)
	internal2Scancode[KeyF7] = sdl.GetScancodeFromKey(sdl.K_F7)
	internal2Scancode[KeyF8] = sdl.GetScancodeFromKey(sdl.K_F8)
	internal2Scancode[KeyF9] = sdl.GetScancodeFromKey(sdl.K_F9)
	internal2Scancode[KeyF10] = sdl.GetScancodeFromKey(sdl.K_F10)
	internal2Scancode[KeyF11] = sdl.GetScancodeFromKey(sdl.K_F11)
	internal2Scancode[KeyF12] = sdl.GetScancodeFromKey(sdl.K_F12)
	internal2Scancode[KeyF13] = sdl.GetScancodeFromKey(sdl.K_F13)
	internal2Scancode[KeyF14] = sdl.GetScancodeFromKey(sdl.K_F14)
	internal2Scancode[KeyF15] = sdl.GetScancodeFromKey(sdl.K_F15)
	internal2Scancode[KeyF16] = sdl.GetScancodeFromKey(sdl.K_F16)
	internal2Scancode[KeyF17] = sdl.GetScancodeFromKey(sdl.K_F17)
	internal2Scancode[KeyF18] = sdl.GetScancodeFromKey(sdl.K_F18)
	internal2Scancode[KeyF19] = sdl.GetScancodeFromKey(sdl.K_F19)
	internal2Scancode[KeyF20] = sdl.GetScancodeFromKey(sdl.K_F20)
	internal2Scancode[KeyF21] = sdl.GetScancodeFromKey(sdl.K_F21)
	internal2Scancode[KeyF22] = sdl.GetScancodeFromKey(sdl.K_F22)
	internal2Scancode[KeyF23] = sdl.GetScancodeFromKey(sdl.K_F23)
	internal2Scancode[KeyF24] = sdl.GetScancodeFromKey(sdl.K_F24)
	internal2Scancode[KeyRight] = sdl.GetScancodeFromKey(sdl.K_RIGHT)
	internal2Scancode[KeyLeft] = sdl.GetScancodeFromKey(sdl.K_LEFT)
	internal2Scancode[KeyDown] = sdl.GetScancodeFromKey(sdl.K_DOWN)
	internal2Scancode[KeyUp] = sdl.GetScancodeFromKey(sdl.K_UP)
	internal2Scancode[KeySpace] = sdl.GetScancodeFromKey(sdl.K_SPACE)
	internal2Scancode[KeyEscape] = sdl.GetScancodeFromKey(sdl.K_ESCAPE)
	internal2Scancode[KeyEnter] = sdl.GetScancodeFromKey(sdl.K_RETURN)
	internal2Scancode[KeyTab] = sdl.GetScancodeFromKey(sdl.K_TAB)
	internal2Scancode[KeyBackspace] = sdl.GetScancodeFromKey(sdl.K_BACKSPACE)
	internal2Scancode[KeyInsert] = sdl.GetScancodeFromKey(sdl.K_INSERT)
	internal2Scancode[KeyDelete] = sdl.GetScancodeFromKey(sdl.K_DELETE)
	internal2Scancode[KeyPageUp] = sdl.GetScancodeFromKey(sdl.K_PAGEUP)
	internal2Scancode[KeyPageDown] = sdl.GetScancodeFromKey(sdl.K_PAGEDOWN)
	internal2Scancode[KeyPrintScreen] = sdl.GetScancodeFromKey(sdl.K_PRINTSCREEN)
	internal2Scancode[KeyPause] = sdl.GetScancodeFromKey(sdl.K_PAUSE)
	internal2Scancode[KeyApostrophe] = sdl.GetScancodeFromKey(sdl.K_QUOTE)
	internal2Scancode[KeyComma] = sdl.GetScancodeFromKey(sdl.K_COMMA)
	internal2Scancode[KeyMinus] = sdl.GetScancodeFromKey(sdl.K_MINUS)
	internal2Scancode[KeyPeriod] = sdl.GetScancodeFromKey(sdl.K_PERIOD)
	internal2Scancode[KeySlash] = sdl.GetScancodeFromKey(sdl.K_SLASH)
	internal2Scancode[KeySemicolon] = sdl.GetScancodeFromKey(sdl.K_SEMICOLON)
	internal2Scancode[KeyEqual] = sdl.GetScancodeFromKey(sdl.K_EQUALS)
	internal2Scancode[KeyLeftBracket] = sdl.GetScancodeFromKey(sdl.K_LEFTBRACKET)
	internal2Scancode[KeyBackslash] = sdl.GetScancodeFromKey(sdl.K_BACKSLASH)
	internal2Scancode[KeyRightBracket] = sdl.GetScancodeFromKey(sdl.K_RIGHTBRACKET)
	internal2Scancode[KeyGraveAccent] = sdl.GetScancodeFromKey(sdl.K_BACKQUOTE)
	internal2Scancode[KeyWWW] = sdl.GetScancodeFromKey(sdl.K_WWW)
	internal2Scancode[KeyHome] = sdl.GetScancodeFromKey(sdl.K_HOME)
	internal2Scancode[KeyEnd] = sdl.GetScancodeFromKey(sdl.K_END)
	internal2Scancode[KeyCapsLock] = sdl.GetScancodeFromKey(sdl.K_CAPSLOCK)
	internal2Scancode[KeyScrollLock] = sdl.GetScancodeFromKey(sdl.K_SCROLLLOCK)
	internal2Scancode[KeyNumLock] = sdl.GetScancodeFromKey(sdl.K_NUMLOCKCLEAR)
	internal2Scancode[KeyNum0] = sdl.GetScancodeFromKey(sdl.K_KP_0)
	internal2Scancode[KeyNum1] = sdl.GetScancodeFromKey(sdl.K_KP_1)
	internal2Scancode[KeyNum2] = sdl.GetScancodeFromKey(sdl.K_KP_2)
	internal2Scancode[KeyNum3] = sdl.GetScancodeFromKey(sdl.K_KP_3)
	internal2Scancode[KeyNum4] = sdl.GetScancodeFromKey(sdl.K_KP_4)
	internal2Scancode[KeyNum5] = sdl.GetScancodeFromKey(sdl.K_KP_5)
	internal2Scancode[KeyNum6] = sdl.GetScancodeFromKey(sdl.K_KP_6)
	internal2Scancode[KeyNum7] = sdl.GetScancodeFromKey(sdl.K_KP_7)
	internal2Scancode[KeyNum8] = sdl.GetScancodeFromKey(sdl.K_KP_8)
	internal2Scancode[KeyNum9] = sdl.GetScancodeFromKey(sdl.K_KP_9)
	internal2Scancode[KeyNumDecimal] = sdl.GetScancodeFromKey(sdl.K_KP_DECIMAL)
	internal2Scancode[KeyNumDivide] = sdl.GetScancodeFromKey(sdl.K_KP_DIVIDE)
	internal2Scancode[KeyNumMultiply] = sdl.GetScancodeFromKey(sdl.K_KP_MULTIPLY)
	internal2Scancode[KeyNumSubtract] = sdl.GetScancodeFromKey(sdl.K_KP_MINUS)
	internal2Scancode[KeyNumAdd] = sdl.GetScancodeFromKey(sdl.K_KP_PLUS)
	internal2Scancode[KeyNumEnter] = sdl.GetScancodeFromKey(sdl.K_KP_ENTER)
	internal2Scancode[KeyNumEqual] = sdl.GetScancodeFromKey(sdl.K_KP_EQUALS)
	internal2Scancode[KeyLeftShift] = sdl.GetScancodeFromKey(sdl.K_LSHIFT)
	internal2Scancode[KeyLeftControl] = sdl.GetScancodeFromKey(sdl.K_LCTRL)
	internal2Scancode[KeyLeftAlt] = sdl.GetScancodeFromKey(sdl.K_LALT)
	internal2Scancode[KeyLeftSuper] = sdl.GetScancodeFromKey(sdl.K_LGUI)
	internal2Scancode[KeyRightShift] = sdl.GetScancodeFromKey(sdl.K_RSHIFT)
	internal2Scancode[KeyRightControl] = sdl.GetScancodeFromKey(sdl.K_RCTRL)
	internal2Scancode[KeyRightAlt] = sdl.GetScancodeFromKey(sdl.K_RALT)
	internal2Scancode[KeyRightSuper] = sdl.GetScancodeFromKey(sdl.K_RGUI)
	internal2Scancode[KeyMenu] = sdl.GetScancodeFromKey(sdl.K_MENU)

	// Create a reverse mapping of sdl scancodes to internal codes for
	// all keys that are handled.
	for code, scan := range internal2Scancode {
		scancode2Internal[scan] = Key(code)
	}

	pretty.Println(internal2Scancode)
	fmt.Println("")
	pretty.Println(scancode2Internal)
}

// KeyPressed returns if key b was just pressed.
func (kb *keyboard) KeyPressed(k Key) bool {
	return kb.input[k] && !kb.lastInput[k]
}

// KeyDown returns if key b is held down.
func (kb *keyboard) KeyDown(k Key) bool {
	return kb.lastInput[k] && kb.input[k]
}

// KeyReleased returns if key b was just released.
func (kb *keyboard) KeyReleased(k Key) bool {
	return kb.lastInput[k] && !kb.input[k]
}

// KeyUp returns if key b is up.
func (kb *keyboard) KeyUp(k Key) bool {
	return !kb.lastInput[k] && !kb.input[k]
}

var (
	// Keyboard exposes all keyboard functionality.
	Keyboard = keyboard{}
)

func (kb *keyboard) updateStatus() {
	kb.any = false
	kb.lastInput = kb.input
}

func processKeyboardEvent(kbe *sdl.KeyboardEvent) {
	inkey := scancode2Internal[kbe.Keysym.Scancode]
	switch kbe.State {
	case sdl.PRESSED:
		Keyboard.any = true
		Keyboard.input[inkey] = true
	case sdl.RELEASED:
		Keyboard.input[inkey] = false
	}
	fmt.Println("Key:", inkey)

	// TODO mod
	// fmt.Println("mod", kbe.Keysym.Mod)
}
