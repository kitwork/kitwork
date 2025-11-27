package main

import (
	"github.com/kitwork/work"
)

func main() {
	err := work.New().
		Secret().
		Schedule().
		Router().
		Run()
	if err != nil {
		panic(err)
	}

}
