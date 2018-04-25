package implementation

import (
	"time"

	"fmt"
	"github.com/HoneyBEP/lua-benchmark/implementation/gopher"
	"github.com/HoneyBEP/lua-benchmark/implementation/native"
	"github.com/HoneyBEP/lua-benchmark/implementation/shopify"
	"io/ioutil"
	"log"
)

// Runs the files in argument on each of the Go-Lua implementations
func Run(files []string) {
	fmt.Print(logTime(native.Run, "NATIVE", files).getString())
	fmt.Print(logTime(gopher.Run, "GOPHER", files).getString())
	fmt.Print(logTime(shopify.Run, "SHOPIFY", files).getString())
}

func Benchmarks() {
	// Read info on files in the benchmarks folder
	fileInfos, err := ioutil.ReadDir("./benchmarks")
	if err != nil {
		log.Panicf("Error reading files: %s", err)
	}

	// Get names of files in 'benchmarks' folder
	files := make([]string, 0, len(fileInfos))
	for _, file := range fileInfos {
		files = append(files, fmt.Sprintf("./benchmarks/%s", file.Name()))
	}

	// Change up order of execution to minimize spin-up advantages
	native1 := logTime(native.Run, "NATIVE", files)
	gopher1 := logTime(gopher.Run, "GOPHER", files)
	shopify1 := logTime(shopify.Run, "SHOPIFY", files)
	shopify2 := logTime(shopify.Run, "SHOPIFY", files)
	gopher2 := logTime(gopher.Run, "GOPHER", files)
	native2 := logTime(native.Run, "NATIVE", files)

	// Tally up the results of the runs
	var nativeRes, gopherRes, shopifyRes *result
	if nativeRes, err = native1.add(native2); err != nil {
		panic(err)
	}
	if gopherRes, err = gopher1.add(gopher2); err != nil {
		panic(err)
	}
	if shopifyRes, err = shopify1.add(shopify2); err != nil {
		panic(err)
	}

	fmt.Printf("\n\n%s%s%s", nativeRes.getString(), gopherRes.getString(), shopifyRes.getString())
}

func Demos() {
	fmt.Println("\nRunning: native demo")
	native.Demo()

	fmt.Println("\nRunning: gopher demo")
	gopher.Demo()

	fmt.Println("\nRunning: shopify demo")
	shopify.Demo()
}

// Measures execution time of the function execution
func logTime(f func([]string) string, name string, tests []string) *result {
	start := time.Now()
	returnString := f(tests)
	elapsed := time.Now().Sub(start) / time.Millisecond

	return &result{name, elapsed, returnString}
}
