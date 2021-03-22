package game

type Vec struct {
	X float32
	Y float32
	Z float32
}

// Add takes another vector and adds to the base vector
// returning the result
func (v1 *Vec) AddVec(vectors ...Vec) Vec {
	for _, v2 := range vectors {
		v1.X += v2.X
		v1.Y += v2.Y
		v1.Z += v2.Z
	}
	return *v1
}

// MultFloat32 takes multiple float32's and multiply's to the base vector
// returning the result
func (v1 *Vec) MultFloat32(vectors ...float32) Vec {
	for _, fl := range vectors {
		v1.X *= fl
		v1.Y *= fl
		v1.Z *= fl
	}
	return *v1
}
