package config

import (
	"os"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v2"
)

type ConfigStruct struct {
}

type ConfigYaml struct {

	//RabbitMQ     RabbitMQStruct      `yaml:"rabbitmq"`

}

// load configuration from env and config.yaml
func Load() (*ConfigStruct, error) {
	// read env variables
	rabbitUser := os.Getenv("RABBIT_USER")
	rabbitPass := os.Getenv("RABBIT_PASS")
	managerUser := os.Getenv("MANAGER_USER")
	managerPass := os.Getenv("MANAGER_PASS")
	databaseUser := os.Getenv("DATABASE_USER")
	databasePass := os.Getenv("DATABASE_PASS")

	// read config.json
	var configYaml = new(ConfigYaml)
	err := defaults.Set(&configYaml)
	if err != nil {
		return nil, err
	}
	file, err := os.ReadFile("./config.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &configYaml)
	if err != nil {
		return nil, err
	}

	//fill config
	var config = new(ConfigStruct)

	return config, nil
}
