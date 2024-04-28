package game

import rl "github.com/gen2brain/raylib-go/raylib"

var cam rl.Camera2D

func createCamera() {

	cam = rl.NewCamera2D(
		rl.NewVector2(player.destination.X+24, player.destination.Y+24),
		rl.NewVector2(player.destination.X+24, player.destination.Y+24),
		0.0,
		2.5)
}

func updateCamera() {

	cam.Target = rl.NewVector2(player.destination.X+24, player.destination.Y+24)

}
