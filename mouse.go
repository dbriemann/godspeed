package godspeed

import "github.com/veandco/go-sdl2/sdl"

type MouseButton uint8

const (
	ButtonLeft MouseButton = iota
	ButtonRight
	ButtonMiddle
	ButtonX1
	ButtonX2
)

type mouse struct {
	pos          Vec2
	lastPos      Vec2
	rel          Vec2
	wheel        Vec2
	buttonStates [5]State
}

func (m *mouse) ButtonState(b MouseButton) State {
	return m.buttonStates[b]
}

var (
	Mouse = mouse{}
)

func updateMouseStatus() {
	for i := 0; i <= int(ButtonX2); i++ {
		switch Mouse.buttonStates[i] {
		case Pressed:
			// If the button was pressed last frame, set it to down.
			// If it receives a new event, the state is overwritten with
			// the correct one.
			Mouse.buttonStates[i] = Down
		case Released:
			// If the button was released last frame, set it to up.
			// If it receives a new event, the state is overwritten with
			// the correct one.
			Mouse.buttonStates[i] = Up
		}
	}
}

func processMouseButtonEvent(mbe *sdl.MouseButtonEvent) {
	state := Up
	switch mbe.State {
	case sdl.PRESSED:
		state = Pressed
	case sdl.RELEASED:
		state = Released
	default:
		// Unknown state.
	}

	switch mbe.Button {
	case sdl.BUTTON_LEFT:
		Mouse.buttonStates[ButtonLeft] = state
	case sdl.BUTTON_RIGHT:
		Mouse.buttonStates[ButtonRight] = state
	case sdl.BUTTON_MIDDLE:
		Mouse.buttonStates[ButtonMiddle] = state
	case sdl.BUTTON_X1:
		Mouse.buttonStates[ButtonX1] = state
	case sdl.BUTTON_X2:
		Mouse.buttonStates[ButtonX2] = state
	default:
		// Unknown button.
	}
}
