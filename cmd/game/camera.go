package game

import rl "github.com/gen2brain/raylib-go/raylib"

var cam rl.Camera2D

func createCamera() {

	cam = rl.NewCamera2D(
		rl.NewVector2(0, 0),
		rl.NewVector2(
			float32(playerDest.X-((SCREENWIDTH/2)-48)),
			float32(playerDest.Y-((SCREENHEIGHT/2)-48)),
		),
		0.0,
		1)
}

func updateCamera() {
	cam.Target = rl.NewVector2(
		float32(playerDest.X-((SCREENWIDTH/2)-48)),
		float32(playerDest.Y-((SCREENHEIGHT/2)-48)),
	)
}
