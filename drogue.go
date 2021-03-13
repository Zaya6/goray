package main

import (
	"drogue/game"
	"drogue/raylib"
	"fmt"
	"runtime"
)

type myBall struct {
	game.Ball
	speed game.Vec
}

func init() {
	runtime.LockOSThread()
}

func main() {
	raylib.InitWindow(800, 600, "raylib in golang!")
	raylib.SetTargetFPS(60)

	testGame := game.Game{}

	var ball = myBall{}
	ball.Pos = game.Vec{X: 400, Y: 300, Z: 0}
	ball.Radius = 50
	ball.Color = raylib.BEIGE
	ball.AddUpdate(func() {

		if raylib.IsKeyDown(raylib.KEY_RIGHT) {
			ball.speed.X += 1
		}
		if raylib.IsKeyDown(raylib.KEY_LEFT) {
			ball.speed.X -= 1
		}
		if raylib.IsKeyDown(raylib.KEY_UP) {
			ball.speed.Y -= 1
		}
		if raylib.IsKeyDown(raylib.KEY_DOWN) {
			ball.speed.Y += 1
		}

		ball.Pos.AddVec(ball.speed)
		ball.speed.MultFloat32(0.9)

	})
	testGame.AddNode(ball.N())

	testGame.StartLoop()

	raylib.CloseWindow()

	fmt.Println("hello drogue!!")
}
