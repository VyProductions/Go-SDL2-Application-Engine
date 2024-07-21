package engine

import "github.com/veandco/go-sdl2/sdl"

type Scene interface {
	// Scene
	Init(*Engine, string) error
	SetActive(bool)
	GetSceneID() string
	Render(*Engine) error
	Free() error

	// Widgets
	InsertWidget(Widget) error
	DeleteWidget(string) error
	GetWidget(string) Widget
	HasWidgetID(string) bool
	GetWidgetIDs() []string

	// Mouse Events
	OnMouseDown(*Engine, sdl.Point) error
	OnMouseUp(*Engine, sdl.Point) error
	OnMouseMove(*Engine, sdl.Point) error
	OnMouseDrag(*Engine, sdl.Point, sdl.Point) error
}
