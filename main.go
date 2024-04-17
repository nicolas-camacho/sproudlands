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

	playerSprite                                  rl.Texture2D
	playerSrc                                     rl.Rectangle
	playerDest                                    rl.Rectangle
	playerMoving                                  bool
	playerDirection                               int
	playerUp, playerDown, playerRight, playerLeft bool
	playerFrame                                   int
	frameCount                                    int

	musicPaused bool
	music       rl.Music

	cam rl.Camera2D
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
		playerMoving = true
		playerDirection = 1
		playerUp = true
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerMoving = true
		playerDirection = 0
		playerDown = true
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerMoving = true
		playerDirection = 2
		playerLeft = true
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerMoving = true
		playerDirection = 3
		playerRight = true
	}

	if rl.IsKeyPressed(rl.KeyQ) {
		musicPaused = !musicPaused
	}
}
func update() {
	running = !rl.WindowShouldClose()

	playerSrc.X = 0

	if playerMoving {
		if playerUp {
			playerDest.Y -= playerSpeed
		}
		if playerDown {
			playerDest.Y += playerSpeed
		}
		if playerRight {
			playerDest.X += playerSpeed
		}
		if playerLeft {
			playerDest.X -= playerSpeed
		}

		if frameCount%8 == 1 {
			playerFrame++
		}

		playerSrc.X = playerSrc.Width * float32(playerFrame)
	}

	frameCount++
	if playerFrame > 3 {
		playerFrame = 0
	}

	playerSrc.Y = playerSrc.Height * float32(playerDirection)

	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}

	cam.Target = rl.NewVector2(
		float32(playerDest.X-(playerDest.Width/2)),
		float32(playerDest.Y-(playerDest.Height/2)),
	)

	playerMoving = false
	playerUp, playerDown, playerRight, playerLeft = false, false, false, false

}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(backgroundColor)
	rl.BeginMode2D(cam)

	drawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}

func quit() {
	rl.UnloadTexture(playerSprite)
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Sproudlands")
	rl.SetExitKey(0)
	rl.SetTargetFPS(fps)

	playerSprite = rl.LoadTexture("resources/Characters/PlayerSpritesheet.png")

	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 100, 100)

	rl.InitAudioDevice()
	music = rl.LoadMusicStream("resources/music/AveryFarm.mp3")
	musicPaused = false
	rl.PlayMusicStream(music)

	cam = rl.NewCamera2D(
		rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
		rl.NewVector2(
			float32(playerDest.X-(playerDest.Width/2)),
			float32(playerDest.Y-(playerDest.Height/2)),
		),
		0.0,
		1.0)
}

func main() {

	for running {
		input()
		update()
		render()
	}
}
