package entities

import (
	"fmt"
	"simple-go-game/internal/assets"
	"simple-go-game/internal/components"
	"simple-go-game/internal/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScoreDisplay_ZIndex    = 1000
	ScoreDisplay_PositionY = 5
)

type ScoreDisplay struct {
	*core.BaseEntity
	*core.BaseDrawable
	value         int
	numberSprites [10]components.Sprite
	numberWidth   float32
	drawArray     []components.Sprite
}

func NewScoreDisplay() *ScoreDisplay {
	sprites := [10]components.Sprite{}
	for i := range sprites {
		sprites[i] = *components.NewSprite(assets.NumberImages[i], components.PivotUpLeft)
	}
	score := ScoreDisplay{
		BaseEntity:    core.NewBaseEntity(),
		BaseDrawable:  core.NewBaseDrawable(ScoreDisplay_ZIndex),
		value:         0,
		numberSprites: sprites,
		numberWidth:   float32(sprites[0].Texture.Width),
	}
	score.calculateDrawArray()
	return &score
}

func (s *ScoreDisplay) Increment() {
	s.value++
	s.calculateDrawArray()
}

func (s *ScoreDisplay) Reset() {
	s.value = 0
	s.calculateDrawArray()
}

func (s *ScoreDisplay) calculateDrawArray() {
	scoreStr := fmt.Sprintf("%d", s.value)
	s.drawArray = s.drawArray[:0] // Clear the slice while retaining capacity

	for _, char := range scoreStr {
		digit := char - '0'
		if digit < 0 || digit > 9 {
			continue // Skip non-digit characters
		}
		s.drawArray = append(s.drawArray, s.numberSprites[digit])
	}
}

func (s *ScoreDisplay) Draw() {
	scoreStr := fmt.Sprintf("%d", s.value)
	totalWidth := float32(len(scoreStr)) * s.numberWidth
	startX := float32(rl.GetScreenWidth()/2) - totalWidth/2

	for i, sprite := range s.drawArray {
		x := startX + float32(i)*s.numberWidth
		sprite.Draw(*core.NewTransform(x, ScoreDisplay_PositionY))
	}
}
