package main

import (
	"github.com/HoneyBEP/lua-benchmark/implementation"
	"os"
	"fmt"
)

func main() {
	args := os.Args[:]
	switch args[1] {
	case "benchmark":
		implementation.Benchmarks()
	case "test":
		implementation.Run(args[2:])
	default:
		fmt.Printf("Argument '%s' is invalid!\n", args[1])
	}
}
