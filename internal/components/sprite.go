package components

import (
	"simple-go-game/internal/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	texture rl.Texture2D
}

// NewSpriteFromBytes crea un sprite desde datos []byte de imagen (PNG)
func NewSpriteFromBytes(data []byte) *Sprite {
	img := rl.LoadImageFromMemory(".png", data, int32(len(data)))
	texture := rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)
	return &Sprite{texture: texture}
}

func (s *Sprite) Draw(transform core.Transform) {
	rl.DrawTexturePro(
		s.texture,
		rl.NewRectangle(0, 0, float32(s.texture.Width), float32(s.texture.Height)), // source
		rl.NewRectangle(
			transform.Position.X,
			transform.Position.Y,
			float32(s.texture.Width)*transform.Scale.X,
			float32(s.texture.Height)*transform.Scale.Y,
		), // dest
		rl.NewVector2(
			float32(s.texture.Width)*transform.Scale.X/2,
			float32(s.texture.Height)*transform.Scale.Y/2,
		), // origin (center for rotation)
		transform.Rotation,
		rl.White,
	)
}
