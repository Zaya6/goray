package raylib

//#include "../c/include/raylib.h"
import "C"

// BeginDrawing raylib wrap
func BeginDrawing() {
	C.BeginDrawing()
}

// EndDrawing raylib wrap
func EndDrawing() {
	C.EndDrawing()
}

// ClearBackground raylib wrap
func ClearBackground() {
	C.ClearBackground(C.RAYWHITE)
}

// DrawText raylib wrap
func DrawText(text string, posX int, posY int, fontSize int) {
	C.DrawText(C.CString(text), C.int(posX), C.int(posY), C.int(fontSize), C.DARKGRAY)
}

// DrawCircleV raylib wrap
func DrawCircleV(centerX float32, centerY float32, radius float32, color Color) {
	//ncolor := C.struct_Color{color.r, color.g, color.b, color.a}
	C.DrawCircle(C.int(centerX), C.int(centerY), C.float(radius), color)
}
