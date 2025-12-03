package main

import (
	"fmt"

	"github.com/kitwork/work"
)

func main() {
	err := work.New().
		Secret().
		Database().
		Schedule().
		Router().
		Run()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}
