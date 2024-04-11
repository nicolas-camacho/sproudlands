package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth          = 1000
	screenHeight         = 480
	fps                  = 60
	playerSpeed  float32 = 3
)

var (
	running         = true
	backgroundColor = rl.NewColor(147, 211, 196, 255)

	playerSprite rl.Texture2D
	playerSrc    rl.Rectangle
	playerDest   rl.Rectangle
)

func drawScene() {
	rl.DrawTexturePro(
		playerSprite,
		playerSrc, playerDest,
		rl.NewVector2(playerDest.Width, playerDest.Height),
		0,
		rl.White)
}

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		playerDest.Y -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerDest.Y += playerSpeed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerDest.X -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerDest.X += playerSpeed
	}
}
func update() {
	running = !rl.WindowShouldClose()
}
func render() {
	rl.BeginDrawing()
	rl.ClearBackground(backgroundColor)

	drawScene()

	rl.EndDrawing()
}

func quit() {
	rl.UnloadTexture(playerSprite)
	rl.CloseWindow()
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Sproudlands")
	rl.SetExitKey(0)
	rl.SetTargetFPS(fps)

	playerSprite = rl.LoadTexture("resources/Characters/PlayerSpritesheet.png")

	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 200, 200)
}

func main() {

	for running {
		input()
		update()
		render()
	}
}
