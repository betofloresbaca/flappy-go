package core

type Pausable interface {
	// IsPaused returns whether the entity is paused.
	IsPaused() bool
	// Pause pauses the entity.
	Pause()
	// Resume resumes the entity.
	Resume()
}
