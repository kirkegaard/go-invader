package core

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Entity interface defines any game object that can be updated and drawn
type Entity interface {
	Update()
	Draw(renderer *sdl.Renderer)
	SetScene(scene *Scene)
}
