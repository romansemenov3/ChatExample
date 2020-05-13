package common

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

var variableRegex = regexp.MustCompile("\\${\\s*(\\w*)(:([^}]*))?}")

// ReadConfig read config.yml to out object
func ReadConfig(out interface{}) error {
	configDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}

	content, err := ioutil.ReadFile(configDir + string(os.PathSeparator) + "config.yml")
	if err != nil {
		return err
	}

	textContent := string(content)
	matches := variableRegex.FindAllStringSubmatch(textContent, -1)
	for _, match := range matches {
		value := ""
		if match[1] != "" {
			val, exists := os.LookupEnv(match[1])
			value = val
			if !exists {
				value = match[3]
			}
		}
		textContent = strings.Replace(textContent, match[0], value, -1)
	}

	return yaml.Unmarshal([]byte(textContent), out)
}
