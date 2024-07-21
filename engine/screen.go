package engine

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
)

type Screen struct {
	// Screen
	sceneID string

	// Scene
	isActive bool
	widgets  map[string]Widget
}

// Scene
func (s *Screen) Init(e *Engine, id string) error {
	if e.HasSceneID(id) {
		return errors.New("scene '" + id + "' already exists in engine")
	}

	s.sceneID = id
	s.isActive = false
	s.widgets = map[string]Widget{}

	return e.InsertScene(s)
}

func (s *Screen) SetActive(active bool) {
	s.isActive = active
}

func (s *Screen) GetSceneID() string {
	return s.sceneID
}

func (s *Screen) Render(e *Engine) error {
	if s.isActive {
		for ID, widget := range s.widgets {
			err := widget.Render(e)

			if err != nil {
				return errors.New(ID + ": " + err.Error())
			}
		}
	}

	return nil
}

func (s *Screen) Free() error {
	for _, widget := range s.widgets {
		if err := widget.Free(); err != nil {
			return err
		}
	}

	return nil
}

// Widgets
func (s *Screen) InsertWidget(w Widget) error {
	id := w.GetWidgetID()

	if s.HasWidgetID(id) {
		return errors.New("widget '" + id + "' already exists")
	}

	s.widgets[id] = w

	return nil
}

func (s *Screen) DeleteWidget(id string) error {
	if !s.HasWidgetID(id) {
		return errors.New("widget '" + id + "' does not exist")
	}

	if err := s.widgets[id].Free(); err != nil {
		return err
	}

	delete(s.widgets, id)

	return nil
}

func (s *Screen) GetWidget(id string) Widget {
	widget, found := s.widgets[id]

	if found {
		return widget
	}

	return nil
}

func (s *Screen) HasWidgetID(id string) bool {
	_, found := s.widgets[id]

	return found
}

func (s *Screen) GetWidgetIDs() []string {
	idList := []string{}

	for id := range s.widgets {
		idList = append(idList, id)
	}

	return idList
}

// Mouse Events
func (s *Screen) OnMouseDown(e *Engine, pos sdl.Point) error {
	if s.isActive {
		for _, widget := range s.widgets {
			if err := widget.OnMouseDown(e, pos); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Screen) OnMouseUp(e *Engine, pos sdl.Point) error {
	if s.isActive {
		for _, widget := range s.widgets {
			if err := widget.OnMouseUp(e, pos); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Screen) OnMouseMove(e *Engine, pos sdl.Point) error {
	if s.isActive {
		for _, widget := range s.widgets {
			if err := widget.OnMouseEnter(e, pos); err != nil {
				return err
			}

			if err := widget.OnMouseLeave(e, pos); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Screen) OnMouseDrag(e *Engine, pos1 sdl.Point, pos2 sdl.Point) error {
	if s.isActive {
		for _, widget := range s.widgets {
			if err := widget.OnMouseDrag(e, pos1, pos2); err != nil {
				return err
			}
		}
	}

	return nil
}
