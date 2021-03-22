// Package display manages the visual representation of the game
package display

import (
	"drogue/raylib"
)

// WindowState is a enum type
type WindowState uint8

const (
	// OPEN WindowState
	OPEN WindowState = iota
	// CLOSED windowState
	CLOSED
	// MINIMIZED windowState
	MINIMIZED
	// MAXIMIZED windowState
	MAXIMIZED
	// FULLSCREEN windowState
	FULLSCREEN
)

// Display is a virtual handle on the raylib window functions
type Display struct {
	Title      string
	Width      int
	Height     int
	Fps        int
	Status     WindowState
	ClearColor raylib.Color
	//handle TODO: Find a way to get the opengl context
}

// Init sets up a raylib window
func (d *Display) Init() *Display {
	raylib.InitWindow(d.Width, d.Height, d.Title)
	raylib.SetTargetFPS(d.Fps)

	for !raylib.IsWindowReady() {
	}
	d.Status = OPEN

	return d
}

// Close destroys the raylib window
//This is necesary to create any new windows, maybe?
func (d *Display) Close() {
	d.Status = CLOSED
	raylib.CloseWindow()
}

// Update checks on the window and updates its state
func (d *Display) Update() {
	if raylib.WindowShouldClose() {
		d.Status = CLOSED
	}
}
