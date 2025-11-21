package action

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Files []string

	Embed bool
}

func (c *Config) Run() error {

	ctx := new(Context)

	for _, file := range c.Files {

		// chỉ xử lý YAML
		ext := filepath.Ext(file)
		if ext != ".yaml" && ext != ".yml" {
			continue
		}

		// 1. đọc file yaml
		workflow, err := readfile(file)
		if err != nil {
			return fmt.Errorf("error reading %s: %w", file, err)
		}

		// 2. parse workflow
		root := Parse(workflow)
		if root == nil {
			return fmt.Errorf("cannot parse workflow: %s", file)
		}

		// 3. run workflow
		fmt.Println(">> Running workflow:", file)

		if err := root.Run(ctx); err != nil {
			return err
		}

	}

	return nil
}

func readfile(path string) (map[string]interface{}, error) {
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

func Source(folder string) *Config {
	cfg := Config{}

	filepath.WalkDir(folder, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if !d.IsDir() {
			cfg.Files = append(cfg.Files, path)
		}

		return nil
	})

	return &cfg
}
