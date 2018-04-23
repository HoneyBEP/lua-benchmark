package implementation

import (
	"time"

	"github.com/HoneyBEP/lua-benchmark/implementation/native"
	"github.com/HoneyBEP/lua-benchmark/implementation/gopher"
	"github.com/HoneyBEP/lua-benchmark/implementation/shopify"
	"fmt"
)

func Run() {
	logTime(native.Run, "NATIVE")

	logTime(gopher.Run, "GOPHER")

	logTime(shopify.Run, "SHOPIFY")
}

func logTime(f func() string, name string) {
	start := time.Now()
	returnString := f()
	elapsed := time.Now().Sub(start)

	fmt.Printf("----- %s BENCHMARK ----- \n Time: %d \n Returned: %s \n Finished \n\n", name, elapsed, returnString)
}
