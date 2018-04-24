package main

import (
	"github.com/HoneyBEP/lua-benchmark/implementation"
	"os"
)

func main() {
	files := os.Args[1:]
	implementation.Run(files)
}
