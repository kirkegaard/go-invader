package main

import (
	"fmt"
	"time"

	"github.com/kirkegaard/go-invader/core"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowWidth  = 800
	windowHeight = 600
	frameRate    = 60
)

func main() {
	// Initialize SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Printf("Failed to initialize SDL: %s\n", err)
		return
	}
	defer sdl.Quit()

	// Create a window
	window, err := sdl.CreateWindow(
		"Go Invader",
		sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		windowWidth, windowHeight,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		fmt.Printf("Failed to create window: %s\n", err)
		return
	}
	defer window.Destroy()

	// Create a renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	renderer.SetLogicalSize(windowWidth, windowHeight)
	renderer.SetScale(1, 1)
	if err != nil {
		fmt.Printf("Failed to create renderer: %s\n", err)
		return
	}
	defer renderer.Destroy()

	// Create an input system
	input := core.NewInputSystem()

	// Create a scene and add entities to it
	scene := &core.Scene{}
	scene.AddEntity(NewPlayer(windowWidth/2, windowHeight/2, input))

	frameDelay := time.Second / frameRate

	running := true

	for running {
		startTime := time.Now()

		// Process SDL events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		// Update scene (which updates all entities)
		scene.Update()

		// Clear the screen
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		// Draw scene (which draws all entities)
		scene.Draw(renderer)

		// Present the screen (swap buffers)
		renderer.Present()

		// Cap the frame rate using time.Sleep
		elapsedTime := time.Since(startTime)
		if remainingTime := frameDelay - elapsedTime; remainingTime > 0 {
			time.Sleep(remainingTime)
		}
	}
}
