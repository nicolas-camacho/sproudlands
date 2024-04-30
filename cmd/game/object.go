package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Object struct {
	texture     rl.Texture2D
	source      rl.Rectangle
	destination rl.Rectangle
	collision   rl.Rectangle
}

func NewObject(texturePath string, x float32, y float32) Object {
	object := Object{
		texture:     rl.LoadTexture(texturePath),
		source:      rl.NewRectangle(0, 0, 48, 48),
		destination: rl.NewRectangle(x, y, 48, 48),
		collision:   rl.NewRectangle(x+16, y+16, 16, 16),
	}
	return object
}
