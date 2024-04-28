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
	PLAYERSPEED  float32 = 2
)

var (
	backgroundColor = rl.NewColor(147, 211, 196, 255)
	frameCount      int
	player          Player

	grassSprite         rl.Texture2D
	hillSprite          rl.Texture2D
	texture             rl.Texture2D
	tileDest            rl.Rectangle
	tileSrc             rl.Rectangle
	tileMap             []int
	srcMap              []string
	mapWidth, mapHeight int

	chestSprite rl.Texture2D
	chestSrc    rl.Rectangle
	chestDest   rl.Rectangle
)

func drawScene() {

	for i, tile := range tileMap {
		tileDest.X = (float32(i%mapWidth) * 16) + (SCREENWIDTH / 2) - float32((mapWidth*16)/2)
		tileDest.Y = (float32(i/mapWidth) * 16) + (SCREENHEIGHT / 2) - float32((mapHeight*16)/2)

		switch srcMap[i] {
		case "h":
			texture = hillSprite
		default:
			texture = grassSprite
		}

		tileSrc.X = tileSrc.Width * float32((tile-1)%int(texture.Width/int32(tileSrc.Width)))
		tileSrc.Y = tileSrc.Height * float32((tile-1)/int(texture.Width/int32(tileSrc.Width)))

		rl.DrawTexturePro(
			texture,
			tileSrc,
			tileDest,
			rl.NewVector2(0, 0),
			0,
			rl.White,
		)
	}

	rl.DrawTexturePro(
		player.texture,
		player.source, player.destination,
		rl.NewVector2(0, 0),
		0,
		rl.White,
	)

	rl.DrawRectangleLines(player.destination.ToInt32().X+16, player.destination.ToInt32().Y+16, 16, 16, rl.Red)

	/*

		rl.DrawTexturePro(
			chestSprite,
			chestSrc,
			chestDest,
			rl.NewVector2(chestDest.Width, chestDest.Height),
			0,
			rl.White,
		)
	*/

	rl.DrawText("Score", int32(player.destination.X)-((SCREENWIDTH+60)/4), int32(player.destination.Y)-((SCREENHEIGHT+60)/4), 8, rl.White)
}

func Input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		player.moving = true
		player.direction = 1
		player.up = true
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		player.moving = true
		player.direction = 0
		player.down = true
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		player.moving = true
		player.direction = 2
		player.left = true
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		player.moving = true
		player.direction = 3
		player.right = true
	}

	music.InputHandler()
}

func Update(running *bool) {
	*running = !rl.WindowShouldClose()

	player.source.X = player.source.Width * float32(player.frame)

	if player.moving {
		if player.up && player.destination.Y > float32((0*mapHeight-1)*16)+SCREENHEIGHT/2-float32(((mapHeight-1)*16)/2)-15 {
			player.Move(Y, -PLAYERSPEED)
		}
		if player.down && player.destination.Y < float32((0*mapHeight-2)*16)+SCREENHEIGHT/2+float32(((mapHeight-2)*16)/2) {
			player.Move(Y, PLAYERSPEED)
		}
		if player.right && (player.destination.X+48) < float32((0*mapWidth)*16)+15+SCREENWIDTH/2+float32(((mapWidth)*16)/2) {
			player.Move(X, PLAYERSPEED)
		}
		if player.left && player.destination.X > float32((0*mapWidth-1)*16)+SCREENWIDTH/2-float32(((mapWidth-1)*16)/2)-6 {
			player.Move(X, -PLAYERSPEED)
		}

		if frameCount%8 == 1 {
			player.frame++
		}
	} else if frameCount%16 == 1 {
		player.frame++
	}

	frameCount++
	if player.frame > 3 {
		player.frame = 0
	}

	if !player.moving && player.frame > 1 {
		player.frame = 0
	}

	player.source.X = player.source.Width * float32(player.frame)
	player.source.Y = player.source.Height * float32(player.direction)

	music.Update()

	updateCamera()

	player.moving = false
	player.up, player.down, player.right, player.left = false, false, false, false

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
	rl.UnloadTexture(player.texture)
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

	player = NewPlayer()

	/*
		chestSprite = rl.LoadTexture("resources/Objects/Chest.png")

		chestSrc = rl.NewRectangle(0, 0, 48, 48)
		chestDest = rl.NewRectangle(16, 16, 48, 48)
	*/

	music.SetInitialValues()

	createCamera()

	loadMap("maps/one.map")
}
