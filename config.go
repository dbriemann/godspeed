package godspeed

type window struct {
	Title      string
	Width      int
	Height     int
	Fullscreen bool
	VSync      bool
}

// Config represents the whole godspeed configuration.
// It is structured into topics and completely public.
var Config = struct {
	Window window
}{
	window{
		Title:      "godspeed game",
		Width:      800,
		Height:     600,
		Fullscreen: false,
		VSync:      false,
	},
}
