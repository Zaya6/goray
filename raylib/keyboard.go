package raylib

//#include "../c/include/raylib.h"
import "C"

const (
	KEY_RIGHT = C.KEY_RIGHT
	KEY_LEFT  = C.KEY_LEFT
	KEY_UP    = C.KEY_UP
	KEY_DOWN  = C.KEY_DOWN
)

// IsKeyDown raylib wrap
func IsKeyDown(keycode C.int) bool {
	if C.IsKeyDown(keycode) {
		return true
	} else {
		return false
	}
}
