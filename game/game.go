package game

import (
	"drogue/display"
	"drogue/raylib"
)

// Game is the base node of any game and tree
type Game struct {
	node
	Display   *display.Display
	paused    bool
	timeScale float32
	doDraw    bool
}

func (g *Game) Draw() {

	raylib.BeginDrawing()
	raylib.ClearBackground(g.Display.ClearColor)
	g.node.Draw()
	raylib.EndDrawing()

}

//have events move up and down the node tree

func (g *Game) Start() {

	if g.Display != nil {
		//if the status is not closed, keep chugging along
		for g.Display.Status != display.CLOSED {

			g.Update(1)
			g.Draw()
			g.Display.Update()
		}

	}

}
