package omniconfig

import "flag"

type FlagConfigurerOption[T comparable] func(*FlagConfigurer[T])

type FlagConfigurer[T comparable] struct {
	Flagset *flag.FlagSet
	Config  *T
	Args    []string
}

func (fc *FlagConfigurer[T]) With(opts ...FlagConfigurerOption[T]) *FlagConfigurer[T] {
	for _, opt := range opts {
		opt(fc)
	}
	return fc
}

func NewFlagConfigurer[T comparable](flagset *flag.FlagSet, config *T, opts ...FlagConfigurerOption[T]) *FlagConfigurer[T] {
	fc := &FlagConfigurer[T]{
		Flagset: flagset,
		Config:  config,
	}

	return fc.With(opts...)
}

func WithFlagConfigurerArgs[T comparable](args []string) FlagConfigurerOption[T] {
	return func(fc *FlagConfigurer[T]) {
		fc.Args = args
	}
}

func (c FlagConfigurer[T]) Load() (*T, error) {
	if err := c.Flagset.Parse(c.Args); err != nil {
		return nil, err
	}
	return c.Config, nil
}
