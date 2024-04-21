package game

import rl "github.com/gen2brain/raylib-go/raylib"

var cam rl.Camera2D

func createCamera() {
	cam = rl.NewCamera2D(
		rl.NewVector2(float32(SCREENWIDTH/2), float32(SCREENHEIGHT/2)),
		rl.NewVector2(
			float32(playerDest.X-(playerDest.Width/2)),
			float32(playerDest.Y-(playerDest.Height/2)),
		),
		0.0,
		1.0)

	cam.Zoom = 2
}

func updateCamera() {
	cam.Target = rl.NewVector2(
		float32(playerDest.X-(playerDest.Width/2)),
		float32(playerDest.Y-(playerDest.Height/2)),
	)
}
