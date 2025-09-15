// Package game provides the core game logic and initialization.
// It handles the main game loop, window management, and scene coordination.
package core

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Game represents the main game instance.
// It manages the game window, scene, and main game loop.
type Game struct {
	scene  *Scene
	width  int32
	height int32
	title  string
	fps    int32
}

// NewGame creates a new game instance with the specified parameters.
func NewGame(width, height int32, title string, fps int32) *Game {
	return &Game{
		scene:  nil,
		width:  width,
		height: height,
		title:  title,
		fps:    fps,
	}
}

// Scene returns the current scene.
func (g *Game) Scene() *Scene {
	return g.scene
}

// SetScene assigns a scene to the game instance.
func (g *Game) SetScene(s *Scene) {
	g.scene = s
}

// Initialize sets up the game window and initializes raylib.
// This should be called before Run().
func (g *Game) Initialize() {
	rl.InitWindow(g.width, g.height, g.title)
	rl.SetTargetFPS(g.fps)
	rl.InitAudioDevice()
}

// Cleanup properly closes the game window and cleans up resources.
// This should be called when the game is shutting down.
func (g *Game) Cleanup() {
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

// Run starts the main game loop.
// This will block until the game window is closed.
func (g *Game) Run() {
	for !rl.WindowShouldClose() {
		deltaTime := rl.GetFrameTime()
		if g.scene != nil {
			// Update
			g.scene.Update(deltaTime)

			// Render
			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
			g.scene.Draw()
			rl.EndDrawing()
		}
	}
}
