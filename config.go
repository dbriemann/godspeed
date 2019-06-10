package godspeed

type win struct {
	Title      string
	Width      int
	Height     int
	Fullscreen bool
	VSync      bool
}

// Config represents the whole godspeed configuration.
// It is structured into topics and completely public.
var Config = struct {
	Window win
}{
	win{
		Title:      "godspeed game",
		Width:      1000,
		Height:     800,
		Fullscreen: false,
		VSync:      false,
	},
}
