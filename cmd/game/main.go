// Main entry point for the simple-go-game.
// This file initializes and starts the game.
package main

import (
	_ "embed"
	"simple-go-game/internal/core"
	"simple-go-game/internal/scenes"
)

func main() {
	// Create a new game instance
	g := core.NewGame(860, 540, "Flappy Go", 60)
	g.Initialize()
	defer g.Cleanup()
	// Create and set the main scene
	g.SetScene(scenes.GameBoard())
	// Start the main game loop
	g.Run()
}
