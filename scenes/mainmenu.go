package scenes

import "main/engine"

type MainMenu struct{}

func (m *MainMenu) Init(e *engine.Engine) error {
	menu := new(engine.Menu)

	menu.Init(e, "MainMenu")

	return nil
}
