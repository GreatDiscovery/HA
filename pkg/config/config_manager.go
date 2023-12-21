package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Configuration struct {
	Debug         bool
	ListenAddress int
	ListenSocket  string
	RaftEnabled   bool
	RaftBind      string
	RaftDataDir   string
	RaftNodes     []string
}

func (c *Configuration) check() error {
	if c.ListenAddress == 0 {
		return errors.New("invalid ListenAddress")
	}
	return nil
}

func NewConfiguration(fileName string) (Configuration, error) {
	config := *newConfiguration()
	if fileName == "" {
		return config, errors.New("empty fileName")
	}
	file, err := os.Open(fileName)
	if err != nil {
		return config, err
	}
	fmt.Printf("start read file %s\n", fileName)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("read file failed", err)
	}
	err = config.check()
	if err != nil {
		return config, err
	}
	return config, nil
}

func newConfiguration() *Configuration {
	return &Configuration{
		Debug:         false,
		ListenAddress: 0,
		ListenSocket:  "",
		RaftEnabled:   false,
	}
}
