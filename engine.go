package godspeed

// Core engine functions which should be overwritten by the user.
var (
	// Init is run once before the game is started.
	Init = func() {

	}

	// Update is called for every logic step of the engine.
	// The parameter dt contains the time in seconds since
	// the last call to Update.
	Update = func(dt float64) {

	}

	// Draw is called for every draw step of the engine.
	// It will use the active RenderTarget to draw onto.
	Draw = func() {

	}
)

// Secondary engine functions that can be ignored by the user most of the time.
var (
	// Loop runs the actual gameloop of a godspeed game.
	Loop = func() {

	}
)
