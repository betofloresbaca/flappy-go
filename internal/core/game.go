// Package game provides the core game logic and initialization.
// It handles the main game loop, window management, and scene coordination.
package core

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

// Game represents the main game instance.
// It manages the game window, scene, and main game loop.
type Game struct {
	root   Entity
	width  int32
	height int32
	title  string
	fps    int32
}

// NewGame creates a new game instance with the specified parameters.
func NewGame(width, height int32, title string, fps int32) *Game {
	return &Game{
		root:   nil,
		width:  width,
		height: height,
		title:  title,
		fps:    fps,
	}
}

// Root returns the current scene.
func (g *Game) Root() Entity {
	return g.root
}

// SetRoot assigns a scene to the game instance.
func (g *Game) SetRoot(e Entity) {
	if g.root != nil {
		g.root.removed()
	}
	g.root = e
	if g.root != nil {
		g.root.added()
	}
}

// Initialize sets up the game window and initializes raylib.
// This should be called before Run().
func (g *Game) Initialize() {
	raylib.InitWindow(g.width, g.height, g.title)
	raylib.SetTargetFPS(g.fps)
	raylib.InitAudioDevice()
}

// Cleanup properly closes the game window and cleans up resources.
// This should be called when the game is shutting down.
func (g *Game) Cleanup() {
	g.SetRoot(nil)
	raylib.CloseAudioDevice()
	raylib.CloseWindow()
}

// Run starts the main game loop.
// This will block until the game window is closed.
func (g *Game) Run() {
	for !raylib.WindowShouldClose() {
		deltaTime := raylib.GetFrameTime()
		if g.root != nil {
			// Update
			if updater, ok := g.root.(Updater); ok {
				updater.Update(deltaTime)
			}

			// Render
			raylib.BeginDrawing()
			raylib.ClearBackground(raylib.RayWhite)
			if drawer, ok := g.root.(Drawer); ok {
				drawer.Draw()
			}
			raylib.EndDrawing()
		}
	}
}
