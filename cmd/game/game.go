package game

import (
	"bufio"
	"fmt"
	"os"
	"sproudlands/internal/music"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREENWIDTH          = 1000
	SCREENHEIGHT         = 480
	FPS                  = 60
	PLAYERSPEED  float32 = 3
)

var (
	backgroundColor = rl.NewColor(147, 211, 196, 255)

	playerSprite                                  rl.Texture2D
	playerSrc                                     rl.Rectangle
	playerDest                                    rl.Rectangle
	playerMoving                                  bool
	playerDirection                               int
	playerUp, playerDown, playerRight, playerLeft bool
	playerFrame                                   int
	frameCount                                    int

	grassSprite         rl.Texture2D
	hillSprite          rl.Texture2D
	texture             rl.Texture2D
	tileDest            rl.Rectangle
	tileSrc             rl.Rectangle
	tileMap             []int
	srcMap              []string
	mapWidth, mapHeight int
)

func drawScene() {
	for i := 0; i < len(tileMap); i++ {
		if tileMap[i] != 0 {
			tileDest.X = tileDest.Width * float32(i%mapWidth)
			tileDest.Y = tileDest.Height * float32(i/mapWidth)

			switch srcMap[i] {
			case "h":
				texture = hillSprite
			default:
				texture = grassSprite
			}

			tileSrc.X = tileSrc.Width * float32((tileMap[i]-1)%int(texture.Width/int32(tileSrc.Width)))
			tileSrc.Y = tileSrc.Height * float32((tileMap[i]-1)/int(texture.Width/int32(tileSrc.Width)))

			rl.DrawTexturePro(
				texture,
				tileSrc, tileDest,
				rl.NewVector2(tileDest.Width, tileDest.Height),
				0,
				rl.White)
		}
	}

	rl.DrawTexturePro(
		playerSprite,
		playerSrc, playerDest,
		rl.NewVector2(playerDest.Width, playerDest.Height),
		0,
		rl.White)
}

func Input() {
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

	music.InputHandler()
}

func Update(running *bool) {
	*running = !rl.WindowShouldClose()

	playerSrc.X = playerSrc.Width * float32(playerFrame)

	if playerMoving {
		if playerUp && playerDest.Y > 16 {
			playerDest.Y -= PLAYERSPEED
		}
		if playerDown && playerDest.Y < 16*(float32(mapHeight)-0.8) {
			playerDest.Y += PLAYERSPEED
		}
		if playerRight && playerDest.X < float32(16*mapWidth) {
			playerDest.X += PLAYERSPEED
		}
		if playerLeft && playerDest.X > 16*1.8 {
			playerDest.X -= PLAYERSPEED
		}

		if frameCount%8 == 1 {
			playerFrame++
		}
	} else if frameCount%16 == 1 {
		playerFrame++
	}

	frameCount++
	if playerFrame > 3 {
		playerFrame = 0
	}

	if !playerMoving && playerFrame > 1 {
		playerFrame = 0
	}

	playerSrc.X = playerSrc.Width * float32(playerFrame)
	playerSrc.Y = playerSrc.Height * float32(playerDirection)

	music.Update()

	updateCamera()

	playerMoving = false
	playerUp, playerDown, playerRight, playerLeft = false, false, false, false

}

func Render() {
	rl.BeginDrawing()
	rl.ClearBackground(backgroundColor)
	rl.BeginMode2D(cam)

	drawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}

func Quit() {
	rl.UnloadTexture(playerSprite)
	music.Unload()
	rl.CloseWindow()
}

func loadMap(mapFile string) {
	file, err := os.Open(mapFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var tileList []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		tileList = append(tileList, fields...)
	}

	mapWidth, mapHeight = -1, -1

	for _, s := range tileList {
		m, err := strconv.Atoi(s)
		if err != nil {
			srcMap = append(srcMap, s)
			continue
		}
		if mapWidth == -1 {
			mapWidth = m
		} else if mapHeight == -1 {
			mapHeight = m
		} else {
			tileMap = append(tileMap, m+1)
		}
	}
	if len(tileMap) > mapWidth*mapHeight {
		tileMap = tileMap[:len(tileMap)-1]
	}
}

func Init() {
	rl.InitWindow(SCREENWIDTH, SCREENHEIGHT, "Sproudlands")
	rl.SetExitKey(0)
	rl.SetTargetFPS(FPS)

	grassSprite = rl.LoadTexture("resources/Tilesets/Grass.png")
	hillSprite = rl.LoadTexture("resources/Tilesets/Hills.png")

	tileSrc = rl.NewRectangle(0, 0, 16, 16)
	tileDest = rl.NewRectangle(0, 0, 16, 16)

	playerSprite = rl.LoadTexture("resources/Characters/PlayerSpritesheetRendererBig.png")

	playerSrc = rl.NewRectangle(0, 0, 192, 192)
	playerDest = rl.NewRectangle(200, 200, 60, 60)

	music.SetInitialValues()

	createCamera()

	loadMap("maps/one.map")
}
