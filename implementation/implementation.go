package implementation

import (
	"time"

	"github.com/HoneyBEP/lua-benchmark/implementation/native"
	"github.com/HoneyBEP/lua-benchmark/implementation/gopher"
	"github.com/HoneyBEP/lua-benchmark/implementation/shopify"
	"fmt"
)

func Run(files []string) {
	logTime(native.Run, "NATIVE", files)

	logTime(gopher.Run, "GOPHER", files)

	logTime(shopify.Run, "SHOPIFY", files)
}

func logTime(f func([]string) string, name string, tests []string) {
	start := time.Now()
	returnString := f(tests)
	elapsed := time.Now().Sub(start) / time.Millisecond

	fmt.Printf("----- %s BENCHMARK ----- \n Time: %d \n Returned: %s \n Finished \n\n", name, elapsed, returnString)
}
