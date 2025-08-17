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
	window           *sdl.Window
	screenSurface    *sdl.Surface
	keyPressSurfaces [KEY_PRESS_SURFACE_TOTAL]*sdl.Surface
)

const (
	KEY_PRESS_SURFACE_DEFAULT int = iota
	KEY_PRESS_SURFACE_UP
	KEY_PRESS_SURFACE_DOWN
	KEY_PRESS_SURFACE_LEFT
	KEY_PRESS_SURFACE_RIGHT
	KEY_PRESS_SURFACE_TOTAL
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
	filepath := "./assets/04/press.bmp"
	keyPressSurfaces[KEY_PRESS_SURFACE_DEFAULT], err = sdl.LoadBMP(filepath)
	if err != nil {
		return fmt.Errorf("Unable to load image %s! SDL Error: %w", filepath, err)
	}

	filepath = "./assets/04/up.bmp"
	keyPressSurfaces[KEY_PRESS_SURFACE_UP], err = sdl.LoadBMP(filepath)
	if err != nil {
		return fmt.Errorf("Unable to load image %s! SDL Error: %w", filepath, err)
	}

	filepath = "./assets/04/down.bmp"
	keyPressSurfaces[KEY_PRESS_SURFACE_DOWN], err = sdl.LoadBMP(filepath)
	if err != nil {
		return fmt.Errorf("Unable to load image %s! SDL Error: %w", filepath, err)
	}

	filepath = "./assets/04/left.bmp"
	keyPressSurfaces[KEY_PRESS_SURFACE_LEFT], err = sdl.LoadBMP(filepath)
	if err != nil {
		return fmt.Errorf("Unable to load image %s! SDL Error: %w", filepath, err)
	}

	filepath = "./assets/04/right.bmp"
	keyPressSurfaces[KEY_PRESS_SURFACE_RIGHT], err = sdl.LoadBMP(filepath)
	if err != nil {
		return fmt.Errorf("Unable to load image %s! SDL Error: %w", filepath, err)
	}

	return nil
}

func closeSDL() {
	for i, v := range keyPressSurfaces {
		v.Free()
		keyPressSurfaces[i] = nil
	}

	_ = window.Destroy()

	sdl.Quit()
}

func loadSurface(path string) (*sdl.Surface, error) {
	loaded, err := sdl.LoadBMP(path)
	if err != nil {
		return nil, fmt.Errorf("On loadBMP `%s`: %w", path, err)
	}
	return loaded, nil
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

	currentSurface := keyPressSurfaces[KEY_PRESS_SURFACE_DEFAULT]
	_ = currentSurface.Blit(nil, screenSurface, nil)
	_ = window.UpdateSurface()

	running := true
	for running {
	eventLoop:
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				running = false
				break eventLoop
			case *sdl.KeyboardEvent:
				if e.Type == sdl.KEYDOWN {
					switch e.Keysym.Sym {
					case sdl.K_UP:
						currentSurface = keyPressSurfaces[KEY_PRESS_SURFACE_UP]
					case sdl.K_DOWN:
						currentSurface = keyPressSurfaces[KEY_PRESS_SURFACE_DOWN]
					case sdl.K_LEFT:
						currentSurface = keyPressSurfaces[KEY_PRESS_SURFACE_LEFT]
					case sdl.K_RIGHT:
						currentSurface = keyPressSurfaces[KEY_PRESS_SURFACE_RIGHT]
					default:
						currentSurface = keyPressSurfaces[KEY_PRESS_SURFACE_DEFAULT]
					}
				}

			}
		}

		_ = currentSurface.Blit(nil, screenSurface, nil)
		_ = window.UpdateSurface()
	}

}
