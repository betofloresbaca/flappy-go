# Simple Go Game

A simple game built with Go and Raylib-go, demonstrating clean architecture and best practices.

## Project Structure

```
simple-go-game/
├── cmd/game/           # Main application entry point
├── internal/           # Private application code
│   ├── core/           # Game engine core systems
│   │   ├── entity/     # Entity system (interfaces and base implementations)
│   │   ├── scene/      # Scene management
│   │   └── game/       # Core game logic and loop
│   └── entities/       # Game-specific entities
│       └── player/     # Player character implementation
├── go.mod              # Go module definition
├── go.sum              # Go module checksums
├── Makefile            # Build and development commands
└── README.md           # This file
```

## Building and Running

To build the game:
```bash
go build -o simple-go-game ./cmd/game
```

To run the game:
```bash
./simple-go-game
```

To run directly without building:
```bash
go run ./cmd/game
```

## Architecture

The project follows Go best practices with a clean architecture:

- **`cmd/game/`**: Contains the main entry point
- **`internal/core/`**: Game engine core systems (reusable)
  - **`entity/`**: Defines the Entity and Drawable interfaces
  - **`scene/`**: Manages collections of entities
  - **`game/`**: Handles the main game loop and window management
- **`internal/entities/`**: Game-specific entity implementations
  - **`player/`**: Player character with movement and rendering

## Dependencies

- [Raylib-go](https://github.com/gen2brain/raylib-go): Go bindings for Raylib game development library

## Development

The code is organized into packages that separate concerns:

1. **Entity System**: Manages game objects with unique IDs and lifecycle methods
2. **Scene Management**: Handles collections of entities with efficient lookup
3. **Rendering**: Supports drawable entities with Z-index sorting
4. **Game Loop**: Manages updates and rendering at 60 FPS