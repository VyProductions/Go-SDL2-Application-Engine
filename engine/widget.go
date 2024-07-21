package engine

import "github.com/veandco/go-sdl2/sdl"

type Widget interface {
	// Widget
	SetActive(bool)
	GetWidgetID() string
	Render(*Engine) error
	Free() error

	// Mouse Events
	OnMouseDown(*Engine, sdl.Point) error
	OnMouseUp(*Engine, sdl.Point) error
	OnMouseEnter(*Engine, sdl.Point) error
	OnMouseLeave(*Engine, sdl.Point) error
	OnMouseDrag(*Engine, sdl.Point, sdl.Point) error
}
