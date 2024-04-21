package main

import "sproudlands/cmd/game"

var running = true

func init() {
	game.Init()
}

func main() {
	defer game.Quit()

	for running {
		game.Input()
		game.Update(&running)
		game.Render()
	}
}
