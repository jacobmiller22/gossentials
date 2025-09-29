package omniconfig

import "io"

type IOConfigurerOption[T any] func(*IOConfigurer[T])

type IOConfigurer[T any] struct {
	R         io.Reader
	Processor func(io.Reader) (*T, error)
}

func (c *IOConfigurer[T]) Load() (*T, error) {
	return c.Processor(c.R)
}

func (c *IOConfigurer[T]) With(opts ...IOConfigurerOption[T]) *IOConfigurer[T] {
	for _, opt := range opts {
		opt(c)
	}
	return c
}
