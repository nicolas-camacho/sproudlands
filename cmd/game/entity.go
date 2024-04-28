package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity struct {
	texture     rl.Texture2D
	source      rl.Rectangle
	destination rl.Rectangle
	collision   rl.Rectangle
	direction   int
	moving      bool
	up          bool
	down        bool
	left        bool
	right       bool
}

type Player struct {
	Entity
}

func (player Player) Move(axis Axis, value int) {

}
