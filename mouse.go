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
	lastPos   Vec2
	pos       Vec2
	rel       Vec2
	wheel     Vec2
	lastInput [5]bool
	input     [5]bool
}

// ButtonPressed returns if button b was just pressed.
func (m *mouse) ButtonPressed(b MouseButton) bool {
	return m.input[b] && !m.lastInput[b]
}

// ButtonDown returns if button b is held down.
func (m *mouse) ButtonDown(b MouseButton) bool {
	return m.lastInput[b] && m.input[b]
}

// ButtonReleased returns if button b was just released.
func (m *mouse) ButtonReleased(b MouseButton) bool {
	return m.lastInput[b] && !m.input[b]
}

// ButtonUp returns if button b is up.
func (m *mouse) ButtonUp(b MouseButton) bool {
	return !m.lastInput[b] && !m.input[b]
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

func (m *mouse) updateStatus() {
	m.rel.X, m.rel.Y = m.pos.X-m.lastPos.X, m.pos.Y-m.lastPos.Y
	m.lastPos.X, m.lastPos.Y = m.pos.X, m.pos.Y

	m.lastInput = m.input
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
	var pressed bool
	switch mbe.State {
	case sdl.PRESSED:
		pressed = true
	case sdl.RELEASED:
		pressed = false
	default:
		// Unknown state.
	}

	switch mbe.Button {
	case sdl.BUTTON_LEFT:
		Mouse.input[ButtonLeft] = pressed
	case sdl.BUTTON_RIGHT:
		Mouse.input[ButtonRight] = pressed
	case sdl.BUTTON_MIDDLE:
		Mouse.input[ButtonMiddle] = pressed
	case sdl.BUTTON_X1:
		Mouse.input[ButtonX1] = pressed
	case sdl.BUTTON_X2:
		Mouse.input[ButtonX2] = pressed
	default:
		// Unknown button.
	}
}
