package game

import (
	"drogue/raylib"
)

// Game is the base node of any game and tree
type Game struct {
	node
	paused    bool
	timeScale float32
}

func (g *Game) Draw() {
	raylib.BeginDrawing()
	raylib.ClearBackground()
	g.node.Draw()
	raylib.EndDrawing()

}

func (g *Game) StartLoop() {
	for !raylib.WindowShouldClose() {
		g.Update(1)
		g.Draw()

	}
}

// N returns the node interface{}
func (n *Game) N() *Node {
	var nd Node = n
	return &nd
}
