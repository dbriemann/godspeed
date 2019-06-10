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

func (m *mouse) Pos() Vec2 {
	return m.pos
}

func (m *mouse) Wheel() Vec2 {
	return m.wheel
}

func (m *mouse) Moved() Vec2 {
	return m.rel
}

var (
	// Mouse exposes all mouse functionality
	Mouse = mouse{}
)

func updateMouseStatus() {
	Mouse.rel.X, Mouse.rel.Y = Mouse.pos.X-Mouse.lastPos.X, Mouse.pos.Y-Mouse.lastPos.Y
	Mouse.lastPos.X, Mouse.lastPos.Y = Mouse.pos.X, Mouse.pos.Y

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

func processMouseMotionEvent(mme *sdl.MouseMotionEvent) {
	Mouse.pos.X = mme.X
	Mouse.pos.Y = mme.Y
	Mouse.rel.X = mme.XRel
	Mouse.rel.Y = mme.YRel
}

func processMouseWheelEvent(mwe *sdl.MouseWheelEvent) {
	mul := int32(1)
	if mwe.Direction == sdl.MOUSEWHEEL_FLIPPED {
		mul = -1
	}

	Mouse.wheel.X = mul * mwe.X
	Mouse.wheel.Y = mul * mwe.Y
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
