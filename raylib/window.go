package raylib

//#include "../c/include/raylib.h"
import "C"

type Window interface {
	InitWindow(int, int, string)
}

type window byte

// InitWindow raylib wrap
func InitWindow(width int, height int, title string) {
	C.InitWindow(C.int(width), C.int(height), C.CString(title))
}

// SetTargetFPS raylib wrap
func SetTargetFPS(targetFPS int) {
	C.SetTargetFPS(C.int(targetFPS))
}

// CloseWindow raylib wrap
func CloseWindow() {
	C.CloseWindow()
}

// WindowShouldClose rayylib wrap
func WindowShouldClose() bool {
	if C.WindowShouldClose() {
		return true
	} else {
		return false
	}
}

func IsWindowReady() bool {
	if C.IsWindowReady() {
		return true
	} else {
		return false
	}
}
