// Main entry point for the flappy-go.
// This file initializes and starts the game.
package main

import (
	_ "embed"
	"flappy-go/internal/core"
	"flappy-go/internal/scenes"
)

func main() {
	// Create a new game instance
	g := core.NewGame(860, 540, "Flappy Go", 60)
	g.Initialize()
	defer g.Cleanup()
	// Create and set the main scene
	g.SetRoot(scenes.GameBoard())
	// Start the main game loop
	g.Run()
}
