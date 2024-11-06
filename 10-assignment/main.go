package main

import (
	"github.com/retry19/challenge-hacktiv8/10-assignment/controllers"
	"github.com/retry19/challenge-hacktiv8/10-assignment/listeners"
)

func main() {
	go controllers.TriggerAutoReload()

	listeners.StartHttpListener()
}
