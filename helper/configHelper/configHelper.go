package confighelper

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func ReadFile(path string, cfg *Config) {
	f, err := os.Open(path)
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func ReadEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}


