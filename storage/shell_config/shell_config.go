package shell_config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/anatolyjudov/job/app/model"
)

type ShellConfigStorage struct{}

func (s ShellConfigStorage) getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".job", "config.json"), nil
}

func (s ShellConfigStorage) Read() (*model.AppConfig, error) {
	configPath, err := s.getConfigFilePath()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config model.AppConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
