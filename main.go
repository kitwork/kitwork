package main

import (
	"github.com/kitwork/engine"
)

func main() {
	err := engine.Run("config.kitwork.yaml")
	if err != nil {
		panic(err)
	}
}
