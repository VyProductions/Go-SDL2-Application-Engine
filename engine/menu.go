package engine

import "errors"

type Menu struct {
	// Menu
	title string

	// Scene
	isActive bool
	widgets  []Widget
}

// Scene
func (m *Menu) Init(e *Engine, title string) error {
	if e.HasSceneID(title) {
		return errors.New("scene '" + title + "' already exists in engine")
	}

	m.title = title
	m.isActive = false
	m.widgets = []Widget{}

	return e.InsertScene(m)
}

func (m *Menu) SetActive(active bool) { m.isActive = active }
func (m *Menu) GetSceneID() string    { return m.title }

// Widgets
func (m *Menu) InsertWidget(Widget) error { return nil }
func (m *Menu) DeleteWidget(string) error { return nil }
func (m *Menu) GetWidget(string) Widget   { return nil }
func (m *Menu) HasWidgetID(string) bool   { return false }
func (m *Menu) GetWidgetIDs() []string    { return nil }
