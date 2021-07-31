package confighelper

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(path string, cfg *Config) (err error) {
	f, err := os.Open(path)
	if err != nil {
		processError(err)
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
	return nil
}
