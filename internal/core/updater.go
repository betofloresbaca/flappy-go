package core

// Updater represents an entity that can be updated every frame.
type Updater interface {
	// Update is called every frame with the delta time in seconds.
	Update(dt float32)
	// IsPaused returns whether the entity is paused.
	Paused() bool
	// Pause pauses the entity.
	Pause()
	// Resume resumes the entity.
	Resume()
}

type BaseUpdater struct {
	paused bool
}

func NewBaseUpdater() *BaseUpdater {
	return &BaseUpdater{paused: false}
}

func (bu *BaseUpdater) Update(dt float32) {}

func (bu *BaseUpdater) Paused() bool {
	return bu.paused
}

func (bu *BaseUpdater) Pause() {
	bu.paused = true
	bu.OnPause()
}

func (bu *BaseUpdater) Resume() {
	bu.paused = false
	bu.OnResume()
}

func (bu *BaseUpdater) OnPause() {}

func (bu *BaseUpdater) OnResume() {}
