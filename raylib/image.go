package raylib

//#include "../c/include/raylib.h"
import "C"

type Image = C.struct_Image
type Texture2d = C.struct_Texture

// GenImagePerlinNoise generates and image with perlin noise
func GenImagePerlinNoise(width int, height int, offsetX int, offsetY int, scale float32) Image {
	return C.GenImagePerlinNoise(C.int(width), C.int(height), C.int(offsetX), C.int(offsetY), C.float(scale))
}

// LoadTextureFromImage turns an image into a Texture2d
func LoadTextureFromImage(image Image) Texture2d {
	return Texture2d(C.LoadTextureFromImage(image))
}
