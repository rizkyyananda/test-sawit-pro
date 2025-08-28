package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Config struct {
	Env    string `yaml:"-"`
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		Port     int    `yaml:"port"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`
}

func LoadConfig(env string) (*Config, error) {
	env = strings.ToLower(strings.TrimSpace(env))
	if env == "" {
		env = "production" // default aman
	}

	var dir string
	var err error
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
		dir, err = os.Getwd()
	} else {
		dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	}
	if err != nil {
		return nil, fmt.Errorf("failed to determine working directory: %w", err)
	}

	filename := map[string]string{
		"local":      "local.yaml",
		"staging":    "staging.yaml",
		"production": "production.yaml",
	}[env]
	if filename == "" {
		filename = "production.yaml"
	}
	configPath := filepath.Join(dir, "env", filename)
	log.Printf("Loading config for env=%q from %s", env, configPath)

	f, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	cfg.Env = env
	return &cfg, nil
}

func InitDB(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password,
		cfg.Database.DBName, cfg.Database.Port, cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	log.Printf("Connected to the database (env=%s)", cfg.Env)
	return db, nil
}
