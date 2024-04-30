package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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
	frame       int
}

type Player struct {
	Entity
}

func (player *Player) Move(axis Axis, speed float32) {
	if axis == X {
		player.destination.X += speed
	}
	if axis == Y {
		player.destination.Y += speed
	}
}

func NewPlayer() Player {
	base := Entity{
		texture:     rl.LoadTexture("resources/Characters/PlayerSpritesheet.png"),
		source:      rl.NewRectangle(0, 0, 48, 48),
		destination: rl.NewRectangle((SCREENWIDTH/2)-24, (SCREENHEIGHT/2)-24, 48, 48),
	}

	base.collision = rl.NewRectangle(base.destination.X+16, base.destination.Y+16, 16, 16)

	return Player{
		base,
	}
}
