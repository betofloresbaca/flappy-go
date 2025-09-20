package core

type BasePausable struct {
	paused bool
}

func NewBasePausable() *BasePausable {
	return &BasePausable{paused: false}
}

func (p *BasePausable) IsPaused() bool {
	return p.paused
}

func (p *BasePausable) Pause() {
	p.paused = true
}

func (p *BasePausable) Resume() {
	p.paused = false
}
