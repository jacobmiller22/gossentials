package omniconfig

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

// const DEFAULT_HISIGHT_CONFIG_FILENAME string = "hisight.json"
// const DEFAULT_HISIGHT_LOGFILE_FILENAME string = "hisight.log"

var ErrMissingConfig error = errors.New("missing config")

// var ErrTooFewArgs error = errors.New("too few args")
// var ErrUnknownLogLevel error = errors.New("unknown log level")

// type key struct{}

// var configKey key

// func WithContext(ctx context.Context, cfg *Config) context.Context {
// 	return context.WithValue(ctx, configKey, cfg)
// }

// func FromContext(ctx context.Context) *Config {
// 	cfg := ctx.Value(configKey)

// 	if cfg, ok := cfg.(*Config); ok && cfg != nil {
// 		return cfg
// 	}

// 	return &DefaultConfig
// }

// func parseLogLevel(s string) (slog.Leveler, error) {
// 	switch strings.ToLower(strings.TrimSpace(s)) {
// 	case "debug":
// 		return slog.LevelDebug, nil
// 	case "info":
// 		return slog.LevelWarn, nil
// 	case "warn", "warning":
// 		return slog.LevelWarn, nil
// 	case "fatal", "error":
// 		return slog.LevelError, nil
// 	default:
// 		return nil, fmt.Errorf("%w: %s", ErrUnknownLogLevel, s)
// 	}
// }

// func parseLogFilepath(s string) (io.Writer, error) {
// 	if s == "stderr" {
// 		return os.Stderr, nil
// 	}
// 	if s == "stdout" {
// 		return os.Stdout, nil
// 	}

// 	if err := os.MkdirAll(filepath.Dir(s), 0644); err != nil {
// 		return nil, err
// 	}

// 	fd, err := os.OpenFile(s, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return fd, nil
// }

// type LogConfig struct {
// 	Level string `json:"level"`
// 	File  string `json:"file"`
// }

// type DBConfig struct {
// 	DSN string `json:"dsn"`
// }

// type HTTPServerConfig struct {
// 	Port int `json:"port"`
// }

// type GRPCServerConfig struct {
// 	Host string `json:"host"`
// 	Port int    `json:"port"`
// }

// type ServerConfig struct {
// 	HTTP HTTPServerConfig `json:"http"`
// 	GRPC GRPCServerConfig `json:"grpc"`
// }

// type Config struct {
// 	Log    LogConfig    `json:"log"`
// 	DB     DBConfig     `json:"db"`
// 	Server ServerConfig `json:"server"`
// 	Args   []string
// }

// func (c *Config) LogLeveler() slog.Leveler {
// 	leveler, err := parseLogLevel(c.Log.Level)
// 	if err != nil {
// 		return slog.LevelError
// 	}
// 	return leveler
// }

// func (c *Config) Logger() (*slog.Logger, error) {
// 	logwriter, err := parseLogFilepath(c.Log.File)
// 	if err != nil {
// 		return nil, err
// 	}

// 	logger := slog.New(slog.NewJSONHandler(
// 		logwriter,
// 		&slog.HandlerOptions{
// 			AddSource: c.LogLeveler() == slog.LevelDebug,
// 			Level:     c.LogLeveler(),
// 		},
// 	))

// 	return logger, nil
//

// var DefaultConfig Config = Config{}

// // func (c *Config) Args() []string {
// // 	if c.args == nil {
// // 		return make([]string, 0)
// // 	}
// // 	return c.args
// // }

// type JsonConfigurer struct {
// 	configPath string
// }

// func DefaultConfigPath() string {
// 	base := os.Getenv("XDG_CONFIG_HOME")
// 	if base == "" {
// 		base = path.Join(os.Getenv("HOME"), ".config")
// 	}
// 	return path.Join(base, DEFAULT_HISIGHT_CONFIG_FILENAME)
// }

// func DefaultLogFilePath() string {
// 	base := os.Getenv("XDG_DATA_HOME")
// 	if base == "" {
// 		base = path.Join(os.Getenv("HOME"), ".local/share")
// 	}
// 	return path.Join(base, DEFAULT_HISIGHT_LOGFILE_FILENAME)
// }

// var DefaulJsonConfigurer = JsonConfigurer{
// 	configPath: DefaultConfigPath(),
// }

func fileExists(p string) bool {
	_, err := os.Stat(p)
	if err != nil {
		return false
	}
	return true
}

// func (c JsonConfigurer) Load() (*Config, error) {
// 	if !fileExists(c.configPath) {
// 		if err := c.Setup(); err != nil {
// 			return nil, err
// 		}
// 	}
// 	fd, err := os.OpenFile(c.configPath, os.O_RDONLY, 0644)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var cfg Config
// 	if err := json.NewDecoder(fd).Decode(&cfg); err != nil {
// 		if errors.Is(err, io.EOF) {
// 			return nil, ErrMissingConfig
// 		}
// 		return nil, err
// 	}

// 	return &cfg, nil
// }

// func (c JsonConfigurer) Setup() error {
// 	fd, err := os.OpenFile(c.configPath, os.O_WRONLY|os.O_CREATE, 0644)
// 	if err != nil {
// 		return err
// 	}

// 	var cfg Config
// 	if err := json.NewEncoder(fd).Encode(cfg); err != nil {
// 		return err
// 	}

// 	return nil
// }

// var DefaultConfigurer = StaticConfigurer{
// 	cfg: &Config{
// 		Log: LogConfig{
// 			Level: "error",
// 			File:  DefaultLogFilePath(),
// 		},
// 		DB: DBConfig{
// 			DSN: ":memory:",
// 		},
// 		Server: ServerConfig{
// 			HTTPServerConfig{
// 				Port: 9001,
// 			},
// 			GRPCServerConfig{
// 				Host: "127.0.0.1:9002",
// 				Port: 9002,
// 			},
// 		},
// 	},
// }

// Configurer defines a way to connect new sources
type Configurer[T comparable] interface {
	Load() (*T, error)
}

// MergeConfigurers() takes variadic number of Configurers and merges them
// together where the priority level rises from the start of the chain to the end.
// Properties where a default value exist will be ignored.
//
// For example:
//
// cfgA = { LogLevel: DEBUG, LogFile: "", LogType: "Unstructured"}
// cfgB = { LogLevel: INFO, LogFile: "./log.log", LogType: ""}
//
// result := MergeConfigurers(cfgA, cfgB)
//
// result is { LogLevel: Info, LogFile: "./log.log", LogType: "Unstructured"}
func MergeConfigurers[T comparable](configurers ...Configurer[T]) (*T, []error, error) {

	var finalCfg *T
	errs := make([]error, len(configurers))

	for i, configurer := range configurers {

		cfg, err := configurer.Load()
		if err != nil {
			errs[i] = err
			continue
		}

		if finalCfg == nil {
			finalCfg = cfg
			continue
		}

		// Perform Merge
		// Type argument "necsesary", without allows invalid pointer arrangement
		if err := mergeStructs[*T](finalCfg, cfg); err != nil {
			return nil, errs, err
		}

	}

	if finalCfg == nil {
		return nil, errs, ErrMissingConfig
	}

	return finalCfg, errs, nil
}

// // return b if b is not the zero value for type T, else a
// func weakAssign[T comparable](a, b T) T {
// 	zero := *new(T)
// 	if b == zero {
// 		return a
// 	}
// 	return b
// }

// mergeStructs merges src into dst using weakAssign rules.
// dst must be a pointer to a struct, src must be a pointer to a struct of the same type.
func mergeStructs[T comparable](dst, src T) error {
	vDst := reflect.ValueOf(dst)
	vSrc := reflect.ValueOf(src)

	if vDst.Kind() != reflect.Ptr || vDst.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("dst must be pointer to struct, got %T", dst)
	}
	if vSrc.Kind() != reflect.Ptr || vSrc.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("src must be pointer to struct, got %T", src)
	}
	if vDst.Type() != vSrc.Type() {
		return fmt.Errorf("dst and src must have same type")
	}

	vDst = vDst.Elem()
	vSrc = vSrc.Elem()

	for i := 0; i < vDst.NumField(); i++ {
		fDst := vDst.Field(i)
		fSrc := vSrc.Field(i)

		// Skip unexported fields
		if !fDst.CanSet() {
			continue
		}

		switch fDst.Kind() {
		case reflect.Struct:
			// Recursively merge nested structs
			err := mergeStructs(fDst.Addr().Interface(), fSrc.Addr().Interface())
			if err != nil {
				return err
			}
		default:
			// Prefer src if it is non-zero
			if !isZero(fSrc) {
				fDst.Set(fSrc)
			}
		}
	}
	return nil
}

// isZero checks whether a reflect.Value is the zero value
func isZero(v reflect.Value) bool {
	zero := reflect.Zero(v.Type()).Interface()
	current := v.Interface()
	return reflect.DeepEqual(current, zero)
}
