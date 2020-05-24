// Package conf provides configuration files in YAML
package configuration

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Configuration is the main struct that represents a configuration.
type Configuration struct {
	Server      string
	Channel     string
	BotName     string
	TLS         bool
	InsecureTLS bool
}

// Config is the Configuration instance that will be exposed to the other packages.
var Config = new(Configuration)

// Load parses the yml file passed as argument and fills the Config.
func Load(cp string) error {
	conf, err := ioutil.ReadFile(cp)
	if err != nil {
		return fmt.Errorf("Conf : Could not read configuration : %v", err)
	}
	if err = yaml.Unmarshal(conf, &Config); err != nil {
		return fmt.Errorf("Conf : Error while parsing yaml : %v", err)
	}
	return nil
}
