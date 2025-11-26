package main

import (
	"github.com/kitwork/work"
)

func main() {
	if err := work.New().
		Secret().
		Schedule().
		Router().
		Run(); err != nil {
		panic(err)
	}

	select {}
	// select {}
}
