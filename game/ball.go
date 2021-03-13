package game

import (
	"drogue/raylib"
)

type Ball struct {
	Circle
	Color raylib.Color
}

func (b *Ball) Update(delta float32) {
	b.node.Update(delta)
}

func (b *Ball) Draw() {
	b.node.Draw()
	raylib.DrawCircleV(b.Pos.X, b.Pos.Y, b.Radius, b.Color)
}

func (n *Ball) N() *Node {
	var nd Node = n
	return &nd
}
