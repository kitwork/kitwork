package main

import (
	"github.com/kitwork/kitwork/core/action"
)

func main() {
	action.Source("./services/tasks").Run()
}
