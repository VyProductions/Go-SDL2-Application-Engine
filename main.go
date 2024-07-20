package main

import (
	"fmt"
	"main/engine"
	"main/scenes"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	// Start Process Engine
	e := new(engine.Engine)

	if err := e.Init(); err != nil {
		panic(err)
	}

	// Initialize Main Menu
	var mainMenu scenes.MainMenu

	if err := mainMenu.Init(e); err != nil {
		panic(err)
	}

	// Process Loop
	running := true

	for running {
		e := sdl.PollEvent()

		if e != nil {
			switch e.(type) {
			case *sdl.KeyboardEvent:
				fmt.Println("Keyboard Event.")
			case *sdl.MouseMotionEvent:
				fmt.Println("MouseMotion Event.")
			case *sdl.MouseButtonEvent:
				fmt.Println("MouseButton Event.")
			case *sdl.MouseWheelEvent:
				fmt.Println("MouseWheel Event.")
			case *sdl.QuitEvent:
				fmt.Println("Quit Event.")
				running = false
			default:
			}
		}
	}

	// Terminate Process Engine
	if err := e.Quit(); err != nil {
		panic(err)
	}
}
