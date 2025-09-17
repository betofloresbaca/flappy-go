package core

type animation struct {
	name         string
	sprites      []Sprite
	frameTime    float32
	loop         bool
	currentFrame int
	elapsedTime  float32
	paused       bool
}

type AnimatedSprite struct {
	animations       map[string]animation
	currentAnimation string
}

func NewAnimatedSprite() *AnimatedSprite {
	return &AnimatedSprite{
		animations: make(map[string]animation),
	}
}

// AddAnimation ahora recibe los frames ([]byte) y los guarda como sprites en Animation
func (as *AnimatedSprite) AddAnimation(name string, frames [][]byte, frameTime float32, loop bool) {
	sprites := make([]Sprite, len(frames))
	for i, data := range frames {
		sprites[i] = *NewSprite(data, PivotCenter)
	}
	as.animations[name] = animation{
		name:      name,
		sprites:   sprites,
		frameTime: frameTime,
		loop:      loop,
	}
}

func (as *AnimatedSprite) SetAnimation(name string) {
	if anim, exists := as.animations[name]; exists {
		as.currentAnimation = name
		anim.currentFrame = 0
		anim.elapsedTime = 0
		anim.paused = false
		as.animations[name] = anim
	}
}

func (as *AnimatedSprite) Play() {
	if anim, exists := as.animations[as.currentAnimation]; exists {
		anim.paused = false
		as.animations[as.currentAnimation] = anim
	}
}

func (as *AnimatedSprite) Pause() {
	if anim, exists := as.animations[as.currentAnimation]; exists {
		anim.paused = true
		as.animations[as.currentAnimation] = anim
	}
}

func (as *AnimatedSprite) Update(dt float32) {
	if as.currentAnimation == "" {
		return
	}
	anim := as.animations[as.currentAnimation]
	if anim.paused {
		return
	}
	anim.elapsedTime += dt
	if anim.elapsedTime >= anim.frameTime {
		anim.elapsedTime = 0
		anim.currentFrame++
		if anim.currentFrame >= len(anim.sprites) {
			if anim.loop {
				anim.currentFrame = 0
			} else {
				anim.currentFrame = len(anim.sprites) - 1 // stay on last frame
				anim.paused = true
			}
		}
	}
	as.animations[as.currentAnimation] = anim
}

func (as *AnimatedSprite) Draw(transform Transform) {
	if as.currentAnimation == "" {
		return
	}
	anim := as.animations[as.currentAnimation]
	if anim.currentFrame >= 0 && anim.currentFrame < len(anim.sprites) {
		anim.sprites[anim.currentFrame].Draw(transform)
	}
}
