package omniconfig

import (
	"encoding/json"
	"io"
)

func JsonReaderProcessor[T any](r io.Reader) (*T, error) {
	decoder := json.NewDecoder(r)
	var cfg T
	err := decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func NewJsonIOConfigurer[T any](r io.Reader) *IOConfigurer[T] {
	return &IOConfigurer[T]{
		R:         r,
		Processor: JsonReaderProcessor[T],
	}
}
