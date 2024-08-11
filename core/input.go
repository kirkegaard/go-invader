package core

import (
	"github.com/veandco/go-sdl2/sdl"
	"sync"
)

type InputEventType int

const (
	KeyDown InputEventType = iota
	KeyUp
)

type InputCallback func(event InputEvent)

type InputEvent struct {
	Type InputEventType
	Key  sdl.Keysym
}

// InputSystem handles all input events
type InputSystem struct {
	callbacks map[InputEventType][]InputCallback
	mu        sync.Mutex
}

func NewInputSystem() *InputSystem {
	return &InputSystem{
		callbacks: make(map[InputEventType][]InputCallback),
	}
}

// On registers a callback for a specific input event
func (is *InputSystem) On(eventType InputEventType, callback InputCallback) {
	is.mu.Lock()
	defer is.mu.Unlock()
	is.callbacks[eventType] = append(is.callbacks[eventType], callback)
}

// ProcessEvent processes an SDL event and triggers callbacks
func (is *InputSystem) ProcessEvent(e sdl.Event) {
	is.mu.Lock()
	defer is.mu.Unlock()

	switch t := e.(type) {
	case *sdl.KeyboardEvent:
		event := InputEvent{
			Key: t.Keysym,
		}
		if t.State == sdl.PRESSED {
			event.Type = KeyDown
		} else if t.State == sdl.RELEASED {
			event.Type = KeyUp
		}

		if callbacks, ok := is.callbacks[event.Type]; ok {
			for _, callback := range callbacks {
				callback(event)
			}
		}
	}
}

// Update can be called in the game loop to process input events
func (is *InputSystem) Update() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		is.ProcessEvent(event)
	}
}
