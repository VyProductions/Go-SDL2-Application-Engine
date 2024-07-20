package engine

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
)

type Engine struct {
	// Display
	window   *sdl.Window
	renderer *sdl.Renderer

	// Scenes
	scenes       map[string]Scene
	currentScene Scene
}

// Setup And Teardown
func (e *Engine) Init() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)

	if err != nil {
		return err
	}

	sdl.SetHint(sdl.HINT_VIDEO_MINIMIZE_ON_FOCUS_LOSS, "0")

	e.window, err = sdl.CreateWindow(
		"Procedural Map Generation",
		0, 0, 1920, 1080,
		sdl.WINDOW_MOUSE_CAPTURE|sdl.WINDOW_MOUSE_FOCUS|
			sdl.WINDOW_FULLSCREEN_DESKTOP,
	)

	if err != nil {
		return err
	}

	e.renderer, err = sdl.CreateRenderer(e.window, -1, sdl.RENDERER_TARGETTEXTURE)

	e.scenes = map[string]Scene{}
	e.currentScene = nil

	return err
}

func (e *Engine) Quit() error {
	err := e.renderer.Destroy()

	if err != nil {
		return err
	}

	err = e.window.Destroy()

	if err != nil {
		return err
	}

	sdl.Quit()

	return nil
}

// Scene Management
func (e *Engine) InsertScene(s Scene) error {
	title := s.GetSceneID()

	if e.HasSceneID(title) {
		return errors.New("scene '" + title + "' already exists")
	}

	e.scenes[title] = s

	return nil
}

func (e *Engine) DeleteScene(title string) error {
	if !e.HasSceneID(title) {
		return errors.New("scene '" + title + "' does not exist")
	}

	delete(e.scenes, title)

	return nil
}

func (e *Engine) GetScene(title string) Scene {
	if !e.HasSceneID(title) {
		return nil
	}

	return e.scenes[title]
}

func (e *Engine) HasSceneID(title string) bool {
	_, foundScene := e.scenes[title]
	return foundScene
}

func (e *Engine) GetSceneIDs() []string {
	idList := []string{}

	for key := range e.scenes {
		idList = append(idList, key)
	}

	return idList
}

func (e *Engine) SwitchScene(title string) error {
	newScene, foundScene := e.scenes[title]

	if !foundScene {
		return errors.New("failed to locate scene '" + title + "'")
	}

	if e.currentScene != nil {
		e.currentScene.SetActive(false)
	}

	e.currentScene = newScene
	e.currentScene.SetActive(true)

	return nil
}
