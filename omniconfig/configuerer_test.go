package omniconfig

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// Test structs

type LogConfig struct {
	Level string `json:"level"`
	File  string `json:"file"`
}

type DBConfig struct {
	DSN string `json:"dsn"`
}

type HTTPServerConfig struct {
	Port int `json:"port"`
}

type GRPCServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ServerConfig struct {
	HTTP HTTPServerConfig `json:"http"`
	GRPC GRPCServerConfig `json:"grpc"`
}

type Config struct {
	Log    LogConfig    `json:"log"`
	DB     DBConfig     `json:"db"`
	Server ServerConfig `json:"server"`
}

func TestMergeConfigurersStatic(t *testing.T) {

	wantConfig := &Config{
		Log: LogConfig{
			Level: "INFO",
		},
		DB: DBConfig{
			DSN: "./tmp.db",
		},
		Server: ServerConfig{
			HTTP: HTTPServerConfig{
				Port: 9001,
			},
			GRPC: GRPCServerConfig{
				Port: 9002,
			},
		},
	}
	wantErrs := []error{nil, nil}

	gotConfig, gotErrs, gotErr := MergeConfigurers(
		StaticConfigurer[Config]{
			Config: &Config{
				Log: LogConfig{
					Level: "DEBUG",
				},
				DB: DBConfig{
					DSN: "./tmp.db",
				},
				Server: ServerConfig{
					HTTP: HTTPServerConfig{
						Port: 9001,
					},
				},
			},
		},
		StaticConfigurer[Config]{
			Config: &Config{
				Log: LogConfig{
					Level: "INFO",
				},
				DB: DBConfig{
					DSN: "",
				},
				Server: ServerConfig{
					GRPC: GRPCServerConfig{
						Port: 9002,
					},
				},
			},
		},
	)
	if gotErr != nil {
		t.Fatalf("Unexpected err: %v", gotErr)
	}

	if diff := cmp.Diff(gotConfig, wantConfig); diff != "" {
		t.Fatalf("mismatch result: %s", diff)
	}

	if diff := cmp.Diff(gotErrs, wantErrs, cmpopts.EquateErrors()); diff != "" {
		t.Fatalf("mismatch errs: %s", diff)
	}
}
