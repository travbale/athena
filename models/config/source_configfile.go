package models

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

type ConfigFileSource struct {
	ConfigurationSource
}

func (config *ConfigFileSource) Config() ([]byte, error) {
	wrapError := func(err error) error {
		return fmt.Errorf("unable to load configuration from file: \n\t%w", err)
	}

	source := config.source

	if !filepath.IsAbs(config.source) {
		currentSource, err := filepath.Abs(config.source)
		if err != nil {
			return nil, wrapError(err)
		}
		source = currentSource
	}

	file, err := os.Open(source)
	if err != nil {
		return nil, wrapError(err)
	}

	defer file.Close()

	bytes := make([]byte, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		bytes = append(bytes, scanner.Bytes()...)
	}

	err = scanner.Err()
	if err != nil {
		return nil, wrapError(err)
	}

	return bytes, nil
}

func (ConfigFileSource) New(source string) IConfigurationSource {
	return &ConfigFileSource{
		ConfigurationSource: ConfigurationSource{
			source_type: "file",
			source:      source,
		},
	}
}
