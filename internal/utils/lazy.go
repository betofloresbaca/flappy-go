package utils

import "sync"

type Lazy[T any] struct {
	value    T
	initFunc func() T
	once     sync.Once
}

func NewLazy[T any](initFunc func() T) *Lazy[T] {
	return &Lazy[T]{initFunc: initFunc}
}

func (l *Lazy[T]) Value() T {
	l.once.Do(func() {
		l.value = l.initFunc()
	})
	return l.value
}
