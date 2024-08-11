package core

import "github.com/veandco/go-sdl2/sdl"

// Scene manages entities and their interactions
type Scene struct {
	entities []Entity
}

func (s *Scene) AddEntity(e Entity) {
	s.entities = append(s.entities, e)
	e.SetScene(s)
}

func (s *Scene) Update() {
	for _, entity := range s.entities {
		entity.Update()
	}
}

func (s *Scene) Draw(renderer *sdl.Renderer) {
	for _, entity := range s.entities {
		entity.Draw(renderer)
	}
}

func (s *Scene) RemoveEntity(e Entity) {
	for i, entity := range s.entities {
		if entity == e {
			s.entities = append(s.entities[:i], s.entities[i+1:]...)
			break
		}
	}
}
