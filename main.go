package main

import (
	"github.com/kitwork/work"
)

func main() {
	err := work.Source("./services/tasks").Run()
	if err != nil {
		panic(err)
	}
	select {}
}
