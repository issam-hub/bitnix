package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App           App
		HTTP          HTTP
		DB            DB
		StorageClient StorageClient
	}

	App struct {
		Name    string `env:"APP_NAME"`
		Version string `env:"VERSION"`
	}

	HTTP struct {
		Port string `env:"HTTP_PORT"`
	}

	DB struct {
		DSN             string `env:"DSN"`
		MaxOpenConns    int    `env:"MAX_OPEN_CONNS"`
		MaxIdleConns    int    `env:"MAX_IDLE_CONNS"`
		MaxIdleLifeTime string `env:"MAX_IDLE_LIFE_TIME"`
	}

	StorageClient struct {
		CloudinaryURL string `env:"CLOUDINARY_URL"`
	}
)

func NewConfig() *Config {
	cfg := &Config{}
	cwd := ProjectRoot()
	envFilePath := cwd + "/.env"
	err := readEnv(envFilePath, cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

func readEnv(envFilePath string, cfg *Config) error {
	envFileExists := checkFileExists(envFilePath)

	if envFileExists {
		err := cleanenv.ReadConfig(envFilePath, cfg)
		if err != nil {
			return fmt.Errorf("config error: %w", err)
		}
	} else {
		err := cleanenv.ReadEnv(cfg)
		if err != nil {
			if _, statErr := os.Stat(envFilePath + ".example"); statErr == nil {
				return fmt.Errorf("missing environment variables: %w\n\nprovide all required environment variables or rename and update .env.example to .env for convinience", err)
			}

			return err
		}
	}
	return nil
}

func checkFileExists(fileName string) bool {
	envFileExists := false
	if _, err := os.Stat(fileName); err == nil {
		envFileExists = true
	}
	return envFileExists
}

func ProjectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(b)

	portions := strings.Split(projectRoot, "/")

	return strings.Join(portions[:len(portions)-1], "/")
}
