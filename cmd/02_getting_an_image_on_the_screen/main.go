package main

import (
	"fmt"

	log "github.com/NikosGour/logging/pkg"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	SCREEN_WIDTH  = 640
	SCREEN_HEIGHT = 480
)

var (
	window        *sdl.Window
	screenSurface *sdl.Surface
	helloWorld    *sdl.Surface
)

func initSDL() error {
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		return fmt.Errorf("SDL could not initialize! SDL_Error: %w", err)
	}

	window, err = sdl.CreateWindow("SDL Tutorial", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, SCREEN_WIDTH, SCREEN_HEIGHT, 0)
	if err != nil {
		return fmt.Errorf("Window could not be created! SDL_Error: %w", err)
	}

	screenSurface, _ = window.GetSurface()
	return nil
}

func loadMedia() error {
	var err error
	filepath := "./assets/02/hello_world.bmp"
	helloWorld, err = sdl.LoadBMP(filepath)
	if err != nil {
		return fmt.Errorf("Unable to load image %s! SDL Error: %w", filepath, err)
	}
	return nil
}

func closeSDL() {
	helloWorld.Free()
	helloWorld = nil

	_ = window.Destroy()

	sdl.Quit()
}
func main() {

	err := initSDL()
	if err != nil {
		log.Fatal("%s", err)
	}
	defer closeSDL()

	err = loadMedia()
	if err != nil {
		log.Fatal("%s", err)
	}

	_ = helloWorld.Blit(nil, screenSurface, nil)
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
