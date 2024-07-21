package engine

import "github.com/veandco/go-sdl2/sdl"

type Image struct {
	// Image
	background   *Rect
	mode         BorderMode
	borderWeight int
	borderColor  sdl.Color
	imageTexture *sdl.Texture
	srcRect      sdl.Rect
	dstRect      sdl.Rect
	draggable    bool

	// Widget
	isActive bool
	widgetID string

	// Mouse Events
}

// Image
func (i *Image) Init(
	id string, imageTexture *sdl.Texture,
	srcRect sdl.Rect, dstRect sdl.Rect,
	mode BorderMode, weight int, draggable bool,
	colors ...sdl.Color,
) error {
	i.widgetID = id
	i.imageTexture = imageTexture
	i.background = new(Rect)
	i.srcRect = srcRect
	i.dstRect = dstRect
	i.mode = mode
	i.borderWeight = weight
	i.draggable = draggable

	if len(colors) > 1 {
		i.borderColor = colors[1]
	} else if len(colors) == 1 {
		i.borderColor = colors[0]
	} else {
		i.borderColor = sdl.Color{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
	}

	i.background.Init(id, dstRect, mode, 0, draggable, colors...)

	return nil
}

// Widget
func (i *Image) SetActive(active bool) {
	i.isActive = active
	i.background.SetActive(active)
}

func (i *Image) GetWidgetID() string {
	return i.widgetID
}

func (i *Image) Render(e *Engine) error {
	if i.isActive {
		// Draw the background
		if err := i.background.Render(e); err != nil {
			return err
		}

		// Draw the image
		if err := e.Renderer.Copy(i.imageTexture, &i.srcRect, &i.dstRect); err != nil {
			return err
		}

		oldR, oldG, oldB, oldA, err := e.Renderer.GetDrawColor()

		if err != nil {
			return err
		}

		// Draw the border
		if err = e.Renderer.SetDrawColor(
			i.borderColor.R, i.borderColor.G, i.borderColor.B, i.borderColor.A,
		); err != nil {
			return err
		}

		var offs int32

		switch i.mode {
		case BORDER_IN:
			offs = 0
		case BORDER_OUT:
			offs = int32(i.borderWeight)
		case BORDER_MID:
			offs = int32(i.borderWeight / 2)
		}

		for j := int32(0); j < int32(i.borderWeight); j++ {
			borderRect := sdl.Rect{
				X: i.dstRect.X - offs + j, Y: i.dstRect.Y - offs + j,
				W: i.dstRect.W + 2*(offs-j), H: i.dstRect.H + 2*(offs-j),
			}

			if err = e.Renderer.DrawRect(&borderRect); err != nil {
				return err
			}
		}

		if err = e.Renderer.SetDrawColor(oldR, oldG, oldB, oldA); err != nil {
			return err
		}
	}

	return nil
}

func (i *Image) Free() error {
	return nil
}

// Mouse Events
func (i *Image) OnMouseDown(e *Engine, pos sdl.Point) error {
	return nil
}

func (i *Image) OnMouseUp(e *Engine, pos sdl.Point) error {
	return nil
}

func (i *Image) OnMouseEnter(e *Engine, pos sdl.Point) error {
	return nil
}

func (i *Image) OnMouseLeave(e *Engine, pos sdl.Point) error {
	return nil
}

func (i *Image) OnMouseDrag(e *Engine, pos1 sdl.Point, pos2 sdl.Point) error {
	if i.draggable {
		return i.background.OnMouseDrag(e, pos1, pos2)
	}

	return nil
}
