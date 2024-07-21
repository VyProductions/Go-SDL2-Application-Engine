package scenes

import (
	"main/engine"

	"github.com/veandco/go-sdl2/sdl"
)

type MainMenu struct{}

func (m *MainMenu) Init(e *engine.Engine) error {
	menu := new(engine.Screen)

	if err := menu.Init(e, "MainMenuScene"); err != nil {
		return err
	}

	startButton := new(engine.Button)
	if err := startButton.Init(
		"StartButton", sdl.Rect{X: 832, Y: 476, W: 256, H: 32}, "Start", engine.NOT_DRAGGABLE,
		func() error {
			return e.SwitchScene("MapScene")
		},
		engine.BORDER_IN, 8, []sdl.Color{
			{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
			{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF},
			{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF},
		}...,
	); err != nil {
		return err
	}

	startButton.SetActive(true)

	if err := menu.InsertWidget(startButton); err != nil {
		return err
	}

	settingsButton := new(engine.Button)
	if err := settingsButton.Init(
		"SettingsButton", sdl.Rect{X: 832, Y: 524, W: 256, H: 32}, "Settings", engine.NOT_DRAGGABLE,
		func() error {
			return e.SwitchScene("SettingsScene")
		},
		engine.BORDER_MID, 8, []sdl.Color{
			{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
			{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF},
			{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF},
		}...,
	); err != nil {
		return err
	}

	settingsButton.SetActive(true)

	if err := menu.InsertWidget(settingsButton); err != nil {
		return err
	}

	quitButton := new(engine.Button)
	if err := quitButton.Init(
		"QuitButton", sdl.Rect{X: 832, Y: 572, W: 256, H: 32}, "Quit", engine.NOT_DRAGGABLE,
		func() error {
			_, err := sdl.PushEvent(&sdl.QuitEvent{
				Type:      sdl.QUIT,
				Timestamp: 0,
			})

			return err
		},
		engine.BORDER_OUT, 8, []sdl.Color{
			{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
			{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF},
			{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF},
		}...,
	); err != nil {
		return err
	}

	quitButton.SetActive(true)

	if err := menu.InsertWidget(quitButton); err != nil {
		return err
	}

	return e.SwitchScene(menu.GetSceneID())
}
