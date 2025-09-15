package entities

import (
	"simple-go-game/internal/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	HSpacing  = 150
	GapHeight = 100
	XStart    = 400
)

type PipeGateGenerator struct {
	*core.BaseEntity
	*core.BaseDrawable
	speed     float32
	pipeGates []PipeGate
	Running   bool
}

func NewPipeGateGenerator(speed float32) *PipeGateGenerator {
	return &PipeGateGenerator{
		BaseEntity:   core.NewBaseEntity(),
		BaseDrawable: core.NewBaseDrawable(-300),
		speed:        speed,
		pipeGates:    make([]PipeGate, 0),
	}
}

func (pg *PipeGateGenerator) addPipe(x float32) {
	y := float32(rl.GetRandomValue(50, int32(int(rl.GetScreenHeight())-int(GapHeight)-100)))
	newPipe := NewPipeGate(x, y, GapHeight, pg.speed)
	newPipe.Running = true
	pg.pipeGates = append(pg.pipeGates, *newPipe)
}

func (pg *PipeGateGenerator) Update(dt float32) {
	if !pg.Running {
		return
	}
	// Update existing pipes
	activePipes := make([]PipeGate, 0, len(pg.pipeGates))
	for i := range pg.pipeGates {
		pg.pipeGates[i].Update(dt)
		if !pg.pipeGates[i].Passed {
			activePipes = append(activePipes, pg.pipeGates[i])
		}
	}
	pg.pipeGates = activePipes

	// Generate new pipes if needed
	if len(pg.pipeGates) == 0 {
		pg.addPipe(XStart)
	}
	nextXStart := pg.getNextXStart()
	for nextXStart < float32(rl.GetScreenWidth()+100) {
		pg.addPipe(nextXStart)
		nextXStart = pg.getNextXStart()
	}

}

func (pg *PipeGateGenerator) getNextXStart() float32 {
	lastPipeX := pg.pipeGates[len(pg.pipeGates)-1].x
	return lastPipeX +
		HSpacing +
		float32(pg.pipeGates[len(pg.pipeGates)-1].topSprite.Texture.Width)
}

func (pg *PipeGateGenerator) Draw() {
	for _, pipe := range pg.pipeGates {
		pipe.Draw()
	}
}
