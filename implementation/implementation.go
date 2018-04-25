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

type result struct {
	name string
	time time.Duration
	result string
}

func (res *result) getString() string {
	return fmt.Sprintf("----- %s BENCHMARK -----\n" +
		"Time: %d\n" +
		"Returned: %s\n" +
		"Finished\n\n", res.name, res.time, res.result)
}

func add(res1 *result, res2 *result) (*result, error) {
	if res1.name != res2.name {
		return nil, fmt.Errorf("methods should be the same, %s and %s are not equal", res1, res2)
	}
	if res1.result != "" || res2.result != "" {
		return nil, fmt.Errorf("one of the methods exited incorrectly")
	}

	return &result{res1.name, res1.time + res2.time, ""}, nil
}

func Run(files []string) {
	fmt.Print(logTime(native.Run, "NATIVE", files).getString())
	fmt.Print(logTime(gopher.Run, "GOPHER", files).getString())
	fmt.Print(logTime(shopify.Run, "SHOPIFY", files).getString())
}

func Benchmarks() {
	fileInfos, err := ioutil.ReadDir("./benchmarks")
	if err != nil {
		log.Panicf("Error reading files: %s", err)
	}

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

	var nativeRes, gopherRes, shopifyRes *result
	if nativeRes, err = add(native1, native2); err != nil {
		panic(err)
	}
	if gopherRes, err = add(gopher1, gopher2); err != nil {
		panic(err)
	}
	if shopifyRes, err = add(shopify1, shopify2); err != nil {
		panic(err)
	}

	fmt.Printf("\n\n%s%s%s", nativeRes.getString(), gopherRes.getString(), shopifyRes.getString())
}

func logTime(f func([]string) string, name string, tests []string) *result {
	start := time.Now()
	returnString := f(tests)
	elapsed := time.Now().Sub(start) / time.Millisecond

	return &result{name, elapsed, returnString}
}
