package omniconfig

type StaticConfigurer[T comparable] struct {
	Config *T
}

func (c StaticConfigurer[T]) Load() (*T, error) {
	return c.Config, nil
}
