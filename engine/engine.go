package engine

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	MOUSE_1 uint8 = uint8(1) // Left Click
	MOUSE_2 uint8 = uint8(3) // Right Click
	MOUSE_3 uint8 = uint8(2) // Middle Click
	MOUSE_4 uint8 = uint8(4) // Back Click
	MOUSE_5 uint8 = uint8(5) // Forward Click

	DRAGGABLE     bool = true
	NOT_DRAGGABLE bool = false
)

type Engine struct {
	// Display
	Window   *sdl.Window
	Renderer *sdl.Renderer

	// Scenes
	scenes       map[string]Scene
	currentScene Scene

	// Input
	mousePos     sdl.Point
	heldKeys     map[sdl.Keycode]bool
	mouseButtons map[uint8]bool
}

// Setup And Teardown
func (e *Engine) Init() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)

	if err != nil {
		return err
	}

	sdl.SetHint(sdl.HINT_VIDEO_MINIMIZE_ON_FOCUS_LOSS, "0")

	e.Window, err = sdl.CreateWindow(
		"Procedural Map Generation",
		0, 0, 1920, 1080,
		sdl.WINDOW_MOUSE_CAPTURE|sdl.WINDOW_MOUSE_FOCUS|
			sdl.WINDOW_INPUT_GRABBED|sdl.WINDOW_FULLSCREEN_DESKTOP,
	)

	if err != nil {
		return err
	}

	e.Renderer, err = sdl.CreateRenderer(e.Window, -1, sdl.RENDERER_TARGETTEXTURE)
	e.Renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)

	e.scenes = map[string]Scene{}
	e.currentScene = nil

	// e.mousePos = sdl.Point{}
	e.mousePos.X, e.mousePos.Y, _ = sdl.GetMouseState()
	e.heldKeys = map[sdl.Keycode]bool{}
	e.mouseButtons = map[uint8]bool{}

	return err
}

func (e *Engine) Quit() error {
	if err := e.Free(); err != nil {
		return err
	}

	if err := e.Renderer.Destroy(); err != nil {
		return err
	}

	if err := e.Window.Destroy(); err != nil {
		return err
	}

	sdl.Quit()

	return nil
}

func (e *Engine) Free() error {
	for _, scene := range e.scenes {
		if err := scene.Free(); err != nil {
			return err
		}
	}

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

// Rendering
func (e *Engine) Render() error {
	if e.currentScene != nil {
		if err := e.Renderer.Clear(); err != nil {
			return err
		}

		if err := e.currentScene.Render(e); err != nil {
			return err
		}

		e.Renderer.Present()
	}

	return nil
}

// Event Handlers
func (e *Engine) HandleKeyboard(event *sdl.KeyboardEvent) error {
	e.heldKeys[event.Keysym.Sym] = event.State == sdl.PRESSED
	return nil
}

func (e *Engine) HandleMouseButton(event *sdl.MouseButtonEvent) error {
	e.mouseButtons[event.Button] = event.State == sdl.PRESSED

	if event.Button == MOUSE_1 {
		if event.State == sdl.PRESSED {
			return e.currentScene.OnMouseDown(e, e.mousePos)
		} else if event.State == sdl.RELEASED {
			return e.currentScene.OnMouseUp(e, e.mousePos)
		}
	}

	return nil
}

func (e *Engine) HandleMouseMotion(event *sdl.MouseMotionEvent) error {
	newPos := sdl.Point{X: event.X, Y: event.Y}

	if e.mouseButtons[MOUSE_1] {
		if err := e.currentScene.OnMouseDrag(e, e.mousePos, newPos); err != nil {
			return err
		}
	}

	e.mousePos.X, e.mousePos.Y = event.X, event.Y

	return e.currentScene.OnMouseMove(e, e.mousePos)
}

func (e *Engine) HandleMouseWheel(event *sdl.MouseWheelEvent) error {
	return nil
}
