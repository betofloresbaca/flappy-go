package scenes

import (
	"simple-go-game/internal/core"
	"simple-go-game/internal/entities/player"
)

func CreateMainBoard() *core.Scene {
	scene := core.NewScene()
	// Add the player to the scene
	player := player.NewPlayer(40, 30, "blue")
	scene.Add(player)
	return scene
}
