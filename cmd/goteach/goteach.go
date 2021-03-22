package main

import (
	"drogue/display"
	"drogue/game"
	"drogue/raylib"
	"fmt"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	dis := display.Display{
		Title:      "GoTeach - v0.0.1",
		Width:      800,
		Height:     600,
		Fps:        60,
		ClearColor: raylib.DARKBLUE,
	}
	dis.Init()

	goteach := game.Game{
		Display: &dis,
	}
	goteach.Start()

	dis.Close()

	fmt.Println("hello teach")

}
