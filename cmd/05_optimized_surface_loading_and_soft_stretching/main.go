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
	window          *sdl.Window
	screenSurface   *sdl.Surface
	strechedSurface *sdl.Surface
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
	filepath := "./assets/05/stretch.bmp"
	strechedSurface, err = sdl.LoadBMP(filepath)
	if err != nil {
		return fmt.Errorf("Unable to load image %s! SDL Error: %w", filepath, err)
	}

	return nil
}

func closeSDL() {
	strechedSurface.Free()
	strechedSurface = nil

	_ = window.Destroy()

	sdl.Quit()
}

func loadSurface(path string) (*sdl.Surface, error) {
	loaded, err := sdl.LoadBMP(path)
	if err != nil {
		return nil, fmt.Errorf("Unable to load image %s! SDL Error: %w", path, err)
	}
	defer loaded.Free()

	optimized, err := loaded.Convert(screenSurface.Format, 0)
	if err != nil {
		return nil, fmt.Errorf("Unable to optimize image %s! SDL Error: %w", path, err)
	}

	return optimized, nil
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

	running := true
	for running {
	eventLoop:
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break eventLoop
			}
		}
		stretchRect := sdl.Rect{}
		stretchRect.X = 0
		stretchRect.Y = 0
		stretchRect.W = SCREEN_WIDTH
		stretchRect.H = SCREEN_HEIGHT
		_ = strechedSurface.BlitScaled(nil, screenSurface, &stretchRect)
		_ = window.UpdateSurface()
	}

}
