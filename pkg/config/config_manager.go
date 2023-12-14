package config

import (
	"HA/pkg/log"
	"context"
	"encoding/json"
	"errors"
	"os"
)

type Configuration struct {
	Debug         bool
	ListenAddress string
	ListenSocket  string
	RaftEnable    bool
}

func NewConfiguration(fileName string) (Configuration, error) {
	config := Configuration{}
	todo := context.TODO()
	if fileName == "" {
		return config, errors.New("empty fileName")
	}
	file, err := os.Open(fileName)
	if err != nil {
		return config, err
	}
	log.G(todo).Infof("start read file %s", fileName)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.G(todo).Error("read file failed", err)
	}
	return config, nil
}
