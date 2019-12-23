package feeder

import (
	"io/ioutil"

	"github.com/goccy/go-yaml"
)

// Yaml is a feeder that feeds using a single json file.
type Yaml struct {
	Path string
}

// Feed returns all the content.
func (j Yaml) Feed() (map[string]interface{}, error) {
	content, err := ioutil.ReadFile(j.Path)
	if err != nil {
		return nil, err
	}

	items := map[string]interface{}{}

	if err := yaml.Unmarshal(content, &items); err != nil {
		return nil, err
	}

	return items, nil
}
