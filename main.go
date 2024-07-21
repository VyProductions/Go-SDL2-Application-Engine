package main

import (
	"main/engine"
	"main/scenes"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	// Start Process Engine
	e := new(engine.Engine)

	// Terminate Process Engine When Exiting
	defer e.Quit()

	if err := e.Init(); err != nil {
		panic(err)
	}

	// Initialize Main Menu Screen
	var mainMenu scenes.MainMenu

	if err := mainMenu.Init(e); err != nil {
		panic(err)
	}

	// Initialize Map Screen
	var mapDisplay scenes.MapDisplay

	if err := mapDisplay.Init(e); err != nil {
		panic(err)
	}

	// Process Loop
	var err error
	running := true

	for running {
		event := sdl.PollEvent()

		if event != nil {
			switch t := event.(type) {
			case *sdl.KeyboardEvent:
				err = e.HandleKeyboard(t)
			case *sdl.MouseMotionEvent:
				err = e.HandleMouseMotion(t)
			case *sdl.MouseButtonEvent:
				err = e.HandleMouseButton(t)
			case *sdl.MouseWheelEvent:
				err = e.HandleMouseWheel(t)
			case *sdl.QuitEvent:
				running = false
			}

			if err != nil {
				panic(err)
			}
		}

		e.Render()
	}
}
