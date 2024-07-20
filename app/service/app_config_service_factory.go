package service

import "github.com/anatolyjudov/job/storage/shell_config"

func AppConfigServiceFactory() *AppConfigService {
	configStorage := shell_config.ShellConfigStorage{}

	return &AppConfigService{
		storage: configStorage,
	}
}
