package main

import (
	"fmt"

	env "github.com/yarles-es/test-initial-golang/config"
)

func main() {
	fmt.Println("Hello World!")
	config := env.GetEnv()
	fmt.Println(config)
}
