package main

import (
	"fmt"
	"github.com/HoneyBEP/lua-benchmark/implementation"
	"os"
)

func main() {
	args := os.Args[:]
	if len(args) < 2 {
		fmt.Printf("Please provide one of the following arguments: '%s', '%s' or '%s'",
			"benchmark", "test", "demo")
		os.Exit(0)
	}

	switch args[1] {
	case "benchmark":
		implementation.Benchmarks()
	case "test":
		implementation.Run(args[2:])
	case "demo":
		implementation.Demos()
	default:
		fmt.Printf("Argument '%s' is invalid!\n", args[1])
	}
}
