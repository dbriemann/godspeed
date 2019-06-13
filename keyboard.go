package godspeed

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Key uint32

func (key Key) String() string {
	return sdl.GetKeyName(sdl.Keycode(key))
}

const (
	KeyUnknown      = Key(sdl.K_UNKNOWN)
	KeySpace        = Key(sdl.K_SPACE)
	KeyApostrophe   = Key(sdl.K_QUOTE)
	KeyComma        = Key(sdl.K_COMMA)
	KeyMinus        = Key(sdl.K_MINUS)
	KeyPeriod       = Key(sdl.K_PERIOD)
	KeySlash        = Key(sdl.K_SLASH)
	Key0            = Key(sdl.K_0)
	Key1            = Key(sdl.K_1)
	Key2            = Key(sdl.K_2)
	Key3            = Key(sdl.K_3)
	Key4            = Key(sdl.K_4)
	Key5            = Key(sdl.K_5)
	Key6            = Key(sdl.K_6)
	Key7            = Key(sdl.K_7)
	Key8            = Key(sdl.K_8)
	Key9            = Key(sdl.K_9)
	KeySemicolon    = Key(sdl.K_SEMICOLON)
	KeyEqual        = Key(sdl.K_EQUALS)
	KeyA            = Key(sdl.K_a)
	KeyB            = Key(sdl.K_b)
	KeyC            = Key(sdl.K_c)
	KeyD            = Key(sdl.K_d)
	KeyE            = Key(sdl.K_e)
	KeyF            = Key(sdl.K_f)
	KeyG            = Key(sdl.K_g)
	KeyH            = Key(sdl.K_h)
	KeyI            = Key(sdl.K_i)
	KeyJ            = Key(sdl.K_j)
	KeyK            = Key(sdl.K_k)
	KeyL            = Key(sdl.K_l)
	KeyM            = Key(sdl.K_m)
	KeyN            = Key(sdl.K_n)
	KeyO            = Key(sdl.K_o)
	KeyP            = Key(sdl.K_p)
	KeyQ            = Key(sdl.K_q)
	KeyR            = Key(sdl.K_r)
	KeyS            = Key(sdl.K_s)
	KeyT            = Key(sdl.K_t)
	KeyU            = Key(sdl.K_u)
	KeyV            = Key(sdl.K_v)
	KeyW            = Key(sdl.K_w)
	KeyX            = Key(sdl.K_x)
	KeyY            = Key(sdl.K_y)
	KeyZ            = Key(sdl.K_z)
	KeyLeftBracket  = Key(sdl.K_LEFTBRACKET)
	KeyBackslash    = Key(sdl.K_BACKSLASH)
	KeyRightBracket = Key(sdl.K_RIGHTBRACKET)
	KeyGraveAccent  = Key(sdl.K_BACKQUOTE)
	KeyWWW          = Key(sdl.K_WWW)
	KeyEscape       = Key(sdl.K_ESCAPE)
	KeyEnter        = Key(sdl.K_RETURN)
	KeyTab          = Key(sdl.K_TAB)
	KeyBackspace    = Key(sdl.K_BACKSPACE)
	KeyInsert       = Key(sdl.K_INSERT)
	KeyDelete       = Key(sdl.K_DELETE)
	KeyRight        = Key(sdl.K_RIGHT)
	KeyLeft         = Key(sdl.K_LEFT)
	KeyDown         = Key(sdl.K_DOWN)
	KeyUp           = Key(sdl.K_UP)
	KeyPageUp       = Key(sdl.K_PAGEUP)
	KeyPageDown     = Key(sdl.K_PAGEDOWN)
	KeyHome         = Key(sdl.K_HOME)
	KeyEnd          = Key(sdl.K_END)
	KeyCapsLock     = Key(sdl.K_CAPSLOCK)
	KeyScrollLock   = Key(sdl.K_SCROLLLOCK)
	KeyNumLock      = Key(sdl.K_NUMLOCKCLEAR)
	KeyPrintScreen  = Key(sdl.K_PRINTSCREEN)
	KeyPause        = Key(sdl.K_PAUSE)
	KeyF1           = Key(sdl.K_F1)
	KeyF2           = Key(sdl.K_F2)
	KeyF3           = Key(sdl.K_F3)
	KeyF4           = Key(sdl.K_F4)
	KeyF5           = Key(sdl.K_F5)
	KeyF6           = Key(sdl.K_F6)
	KeyF7           = Key(sdl.K_F7)
	KeyF8           = Key(sdl.K_F8)
	KeyF9           = Key(sdl.K_F9)
	KeyF10          = Key(sdl.K_F10)
	KeyF11          = Key(sdl.K_F11)
	KeyF12          = Key(sdl.K_F12)
	KeyF13          = Key(sdl.K_F13)
	KeyF14          = Key(sdl.K_F14)
	KeyF15          = Key(sdl.K_F15)
	KeyF16          = Key(sdl.K_F16)
	KeyF17          = Key(sdl.K_F17)
	KeyF18          = Key(sdl.K_F18)
	KeyF19          = Key(sdl.K_F19)
	KeyF20          = Key(sdl.K_F20)
	KeyF21          = Key(sdl.K_F21)
	KeyF22          = Key(sdl.K_F22)
	KeyF23          = Key(sdl.K_F23)
	KeyF24          = Key(sdl.K_F24)
	KeyNum0         = Key(sdl.K_KP_0)
	KeyNum1         = Key(sdl.K_KP_1)
	KeyNum2         = Key(sdl.K_KP_2)
	KeyNum3         = Key(sdl.K_KP_3)
	KeyNum4         = Key(sdl.K_KP_4)
	KeyNum5         = Key(sdl.K_KP_5)
	KeyNum6         = Key(sdl.K_KP_6)
	KeyNum7         = Key(sdl.K_KP_7)
	KeyNum8         = Key(sdl.K_KP_8)
	KeyNum9         = Key(sdl.K_KP_9)
	KeyNumDecimal   = Key(sdl.K_KP_DECIMAL)
	KeyNumDivide    = Key(sdl.K_KP_DIVIDE)
	KeyNumMultiply  = Key(sdl.K_KP_MULTIPLY)
	KeyNumSubtract  = Key(sdl.K_KP_MINUS)
	KeyNumAdd       = Key(sdl.K_KP_PLUS)
	KeyNumEnter     = Key(sdl.K_KP_ENTER)
	KeyNumEqual     = Key(sdl.K_KP_EQUALS)
	KeyLeftShift    = Key(sdl.K_LSHIFT)
	KeyLeftControl  = Key(sdl.K_LCTRL)
	KeyLeftAlt      = Key(sdl.K_LALT)
	KeyLeftSuper    = Key(sdl.K_LGUI)
	KeyRightShift   = Key(sdl.K_RSHIFT)
	KeyRightControl = Key(sdl.K_RCTRL)
	KeyRightAlt     = Key(sdl.K_RALT)
	KeyRightSuper   = Key(sdl.K_RGUI)
	KeyMenu         = Key(sdl.K_MENU)
	KeyLast         = Key(0)
)

type keyboard struct {
	lastInput [sdl.NUM_SCANCODES]bool
	input     [sdl.NUM_SCANCODES]bool
	any       bool
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

// TODO get key descriptor

var (
	// Keyboard exposes all keyboard functionality.
	Keyboard = keyboard{}
)

func (kb *keyboard) updateStatus() {
	kb.any = false
	kb.lastInput = kb.input
}

func processKeyboardEvent(kbe *sdl.KeyboardEvent) {
	switch kbe.State {
	case sdl.PRESSED:
		Keyboard.any = true
		Keyboard.input[kbe.Keysym.Scancode] = true
	case sdl.RELEASED:
		Keyboard.input[kbe.Keysym.Scancode] = false
	}
	fmt.Println("Key:", sdl.GetKeyName(kbe.Keysym.Sym))

	// TODO mod
	// fmt.Println("mod", kbe.Keysym.Mod)
}
