package core

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	Sprite_Format = ".png"
)

type Sprite struct {
	Texture raylib.Texture2D
	Pivot   Pivot
	FlipH   bool
	FlipV   bool
}

// NewSprite creates a sprite from image data ([]byte, PNG) and a pivot
func NewSprite(data []byte, pivot Pivot) *Sprite {
	img := raylib.LoadImageFromMemory(Sprite_Format, data, int32(len(data)))
	texture := raylib.LoadTextureFromImage(img)
	raylib.UnloadImage(img)
	return &Sprite{
		Texture: texture,
		Pivot:   pivot,
	}
}

func (s *Sprite) Draw(transform Transform) {
	width := float32(s.Texture.Width) * transform.Scale.X
	height := float32(s.Texture.Height) * transform.Scale.Y
	var origin raylib.Vector2
	switch s.Pivot {
	case PivotUpLeft:
		origin = raylib.NewVector2(0, 0)
	case PivotUpRight:
		origin = raylib.NewVector2(width, 0)
	case PivotDownLeft:
		origin = raylib.NewVector2(0, height)
	case PivotDownRight:
		origin = raylib.NewVector2(width, height)
	case PivotCenter:
		origin = raylib.NewVector2(width/2, height/2)
	}
	// Calculate source rectangle considering FlipH and FlipV
	srcX := float32(0)
	srcY := float32(0)
	srcW := float32(s.Texture.Width)
	srcH := float32(s.Texture.Height)
	if s.FlipH {
		srcW = -srcW
		srcX = float32(s.Texture.Width)
	}
	if s.FlipV {
		srcH = -srcH
		srcY = float32(s.Texture.Height)
	}
	raylib.DrawTexturePro(
		s.Texture,
		raylib.NewRectangle(srcX, srcY, srcW, srcH), // source
		raylib.NewRectangle(
			transform.Position.X,
			transform.Position.Y,
			width,
			height,
		), // dest
		origin,
		transform.Rotation,
		raylib.White,
	)
}
