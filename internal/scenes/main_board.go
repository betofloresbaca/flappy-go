package scenes

import (
	"simple-go-game/internal/core/scene"
	"simple-go-game/internal/entities/player"
)

func CreateMainBoard() *scene.Scene {
	scene := scene.NewScene()
	// Add the player to the scene
	player := player.NewPlayer(40, 30)
	scene.Add(player)
	return scene
}
