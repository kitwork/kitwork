package handler

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func Readfile(path string) (map[string]interface{}, error) {
	// 1. Đọc file YAML
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %w", err)
	}

	// 2. Parse YAML vào map[string]interface{}
	var workflow map[string]interface{}
	err = yaml.Unmarshal(data, &workflow)
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML: %w", err)
	}

	return workflow, nil
}
