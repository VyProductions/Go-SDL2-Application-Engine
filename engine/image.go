package engine

import "github.com/veandco/go-sdl2/sdl"

type Image struct {
	// Image
	background   Rect
	imageSurface *sdl.Surface

	// Widget
	isActive bool
	widgetID string

	// Mouse Events
}

func (i *Image) SetActive(active bool) { i.isActive = active }
func (i *Image) GetWidgetID() string   { return i.widgetID }

func (i *Image) OnMouseDown(*Engine, sdl.Point)            {}
func (i *Image) OnMouseUp(*Engine, sdl.Point)              {}
func (i *Image) OnMouseEnter(*Engine, sdl.Point)           {}
func (i *Image) OnMouseLeave(*Engine, sdl.Point)           {}
func (i *Image) OnMouseDrag(*Engine, sdl.Point, sdl.Point) {}
