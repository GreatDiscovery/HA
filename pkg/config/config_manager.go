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
	LogRetain     int
	MySQLHostName string
	MySQLPort     int
	MySQLUser     string
	MySQLPassword string
	MySQLDB       string
}

func (c *Configuration) checkAndAssign() error {
	if c.ListenAddress == 0 {
		return errors.New("invalid ListenAddress")
	}
	if c.LogRetain == 0 {
		c.LogRetain = 10
	}
	if c.MySQLHostName == "" {
		c.MySQLHostName = "127.0.0.1"
	}
	if c.MySQLPort == 0 {
		c.MySQLPort = 3306
	}
	if c.MySQLUser == "" {
		c.MySQLUser = "root"
	}
	if c.MySQLPassword == "" {
		return errors.New("empty MySQLPassword")
	}
	if c.MySQLDB == "" {
		return errors.New("empty MySQLDB")
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
	err = config.checkAndAssign()
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
