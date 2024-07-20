package engine

import "github.com/veandco/go-sdl2/sdl"

type Widget interface {
	// Widget
	SetActive(bool)
	GetWidgetID() string

	// Mouse Events
	OnMouseDown(*Engine, sdl.Point)
	OnMouseUp(*Engine, sdl.Point)
	OnMouseEnter(*Engine, sdl.Point)
	OnMouseLeave(*Engine, sdl.Point)
	OnMouseDrag(*Engine, sdl.Point, sdl.Point)
}
