package omniconfig

import "flag"

type FlagConfigurerOption[T comparable] func(*FlagConfigurer[T])

type FlagConfigurer[T comparable] struct {
	Flagset *flag.FlagSet
	Args    []string
}

func NewFlagConfigurer[T comparable](flagset *flag.FlagSet, opts ...FlagConfigurerOption[T]) *FlagConfigurer[T] {
	return &FlagConfigurer[T]{
		Flagset: flagset,
	}
}

func WithFlagConfigurerArgs[T comparable](args []string) FlagConfigurerOption[T] {
	return func(fc *FlagConfigurer[T]) {
		fc.Args = args
	}
}

func (c FlagConfigurer[T]) Load() (*T, error) {
	var cfg T

	// c.flagset.IntVar(&cfg.Server.HTTP.Port, "server-http-port", 0, "Port to listen on for HTTP requests")
	// c.flagset.StringVar(&cfg.Server.GRPC.Host, "server-grpc-host", "", "Hostname of the GRPC Server")
	// c.flagset.IntVar(&cfg.Server.GRPC.Port, "server-grpc-port", 0, "Port to listen on for GRPC requests")
	// c.flagset.StringVar(&cfg.DB.DSN, "db-dsn", "", "DSN to database")
	// c.flagset.StringVar(&cfg.Log.Level, "log-level", "", "Log level to use")
	// c.flagset.StringVar(&cfg.Log.File, "log-file", "", "Log file to use")

	if err := c.Flagset.Parse(c.Args); err != nil {
		return nil, err
	}

	return &cfg, nil
}
