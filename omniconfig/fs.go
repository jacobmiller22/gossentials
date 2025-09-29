package omniconfig

import (
	"os"
)

func NewFsIOConfigurer[T any](path string) (*IOConfigurer[T], error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return &IOConfigurer[T]{
		R: fd,
	}, nil
}
