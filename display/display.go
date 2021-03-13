// Package display manages the visual representation of the game
package display

// WindowState is a enum type
type WindowState uint8

const (
	// OPEN WindowState
	OPEN int = iota
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
	Width  int
	Height int
	Fps    int
	Status WindowState
	//handle TODO: Find a way to get the opengl context
}

// Init sets up a raylib window
func (d *Display) Init() {
	/*
		C.InitWindow(C.int(d.Width), C.int(d.Height), C.CString("hello world"))
		C.SetTargetFPS(C.int(d.Fps))
	*/
	d.Status = WindowState(OPEN)
}

// Update draws everything to screen
func (d *Display) Update() {
	/*
		for !C.WindowShouldClose() {
			if C.IsKeyDown(C.KEY_RIGHT) {
				ball.x += C.float(2.0)
			}
			if C.IsKeyDown(C.KEY_LEFT) {
				ball.x -= C.float(2.0)
			}
			if C.IsKeyDown(C.KEY_UP) {
				ball.y -= C.float(2.0)
			}
			if C.IsKeyDown(C.KEY_DOWN) {
				ball.y += C.float(2.0)
			}

			C.BeginDrawing()

			C.ClearBackground(C.RAYWHITE)

			C.DrawText(C.CString("move the ball with arrow keys"), C.int(10), C.int(10), C.int(20), C.DARKGRAY)

			C.DrawCircleV(ball, C.float(50), C.MAROON)

			C.EndDrawing()
		}

		d.Close()
	*/
}

// Close destroys the raylib window
//This is necesary to create any new windows, maybe?
func (d *Display) Close() {
	/*
		d.Status = WindowState(CLOSED)
		C.CloseWindow()
	*/
}
