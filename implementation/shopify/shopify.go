package shopify

import (
	"fmt"
	"github.com/Shopify/go-lua"
)

func Run(tests []string) string {
	L := lua.NewState()
	lua.OpenLibraries(L)

	lua.DoFile(L, "benchmark.lua")
	if err := lua.DoFile(L, "benchmark.lua"); err != nil {
		return fmt.Sprintf("%s \n", err)
	}

	for _, test := range tests {
		if err := lua.DoFile(L, test); err != nil {
			return fmt.Sprintf("%s \n", err)
		}
	}

	// Lua function in Go
	// lua.DoString(L,"print(benchmark(10))")
	L.Global("benchmark")
	L.PushInteger(10)

	L.Call(1, 1)
	L.ToInteger(1)
	fmt.Printf("print `depth` from Go: %v \n",  L.ToValue(1))

	// Register Go function in Lua
	L.Register("count", func(l *lua.State) int {
		L.PushInteger(8)
		fmt.Println("echo `count` in Go")
		return 1
	})

	// Run Go function in Lua
	// L.Call(0, 1)
	lua.DoString(L, "countFromGo()")

	return ""
}


