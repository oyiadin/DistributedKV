package internal

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type Config struct {
	Peers []struct {
		Name string
		URL  string `yaml:"url"`
	}
	PeerName string
}

func loadConfigWithFilename(fromDirectories []string, filename string) (map[interface{}]interface{}, error) {
	conf := map[interface{}]interface{}{}
	for _, dir := range fromDirectories {
		filePath := path.Join(dir, filename)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			log.Printf("warning: config file %s not exists", filePath)
			continue
		}

		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("failed to read config file %s: %v", filePath, err))
		}
		err = yaml.Unmarshal(data, &conf)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("failed to unmarshal config file %s: %v", filePath, err))

		}

		return conf, nil
	}

	return nil, errors.New(fmt.Sprintf(
		"no any %s has been found within following directories: %v", filename, fromDirectories))
}

func loadDefaultConfig(fromDirectories []string) (conf map[interface{}]interface{}, err error) {
	return loadConfigWithFilename(fromDirectories, "config-default.yaml")
}

func (c *Config) LoadFrom(fromDirectories []string) error {
	defaultConfigMap, err := loadDefaultConfig(fromDirectories)
	if err != nil {
		return err
	}

	confMap, err := loadConfigWithFilename(fromDirectories, "config.yaml")
	if err != nil {
		return err
	}

	mergedConfMap := mergeConfigs(defaultConfigMap, confMap).(map[interface{}]interface{})

	err = mapstructure.Decode(mergedConfMap, c)
	if err != nil {
		return err
	}

	log.Println("succeeded to load config!")
	return err
}

func mergeConfigs(base, updates interface{}) interface{} {
	b, ok1 := base.(map[interface{}]interface{})
	u, ok2 := updates.(map[interface{}]interface{})
	if ok1 && ok2 {
		for key, value := range u {
			if _, ok := b[key]; !ok {
				b[key] = value
			} else {
				b[key] = mergeConfigs(b[key], value)
			}
		}
		return b
	} else {
		return updates
	}
}
