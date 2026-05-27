package main

import (
	"github.com/kitwork/engine"
)

func main() {
	err := engine.Run("config.kitwork.yml")
	if err != nil {
		panic(err)
	}
}
