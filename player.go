package main

import (
	"github.com/kirkegaard/go-invader/core"
	"github.com/veandco/go-sdl2/sdl"
)

type PlayerEntity struct {
	x, y  int32
	input *core.InputSystem
	scene *core.Scene
}

func NewPlayer(x, y int32, input *core.InputSystem) *PlayerEntity {
	player := &PlayerEntity{x: x, y: y, input: input}

	// Register input callbacks
	input.On(core.KeyDown, func(event core.InputEvent) {
		switch event.Key.Sym {
		// case sdl.K_SPACE:
		//     player.EmitBullet()
		case sdl.K_LEFT:
			player.x = -5
		case sdl.K_RIGHT:
			player.x = 5
		}
	})

	return player
}

func (p *PlayerEntity) SetScene(scene *core.Scene) {
	p.scene = scene
}

func (p *PlayerEntity) Update() {
}

func (p *PlayerEntity) Draw(renderer *sdl.Renderer) {
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.FillRect(&sdl.Rect{X: p.x, Y: p.y, W: 10, H: 10})
}
