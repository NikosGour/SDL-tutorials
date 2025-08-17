package main

import (
	log "github.com/NikosGour/logging/pkg"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	SCREEN_WIDTH  = 640
	SCREEN_HEIGHT = 480
)

func main() {
	var window *sdl.Window
	var screenSurface *sdl.Surface

	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		log.Fatal("SDL could not initialize! SDL_Error: %s", err)
	}
	defer sdl.Quit()

	window, err = sdl.CreateWindow("SDL Tutorial", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, SCREEN_WIDTH, SCREEN_HEIGHT, 0)
	if err != nil {
		log.Fatal("Window could not be created! SDL_Error: %s", err)
	}
	defer func() { _ = window.Destroy() }()

	_ = window.SetFullscreen(0)

	screenSurface, _ = window.GetSurface()
	_ = screenSurface.FillRect(nil, sdl.MapRGB(screenSurface.Format, 0xff, 0xff, 0xff))

	_ = window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}
	}

}
