package gopher

import (
	"fmt"
	"github.com/yuin/gopher-lua"
)

func Run(tests []string) string {
	L := lua.NewState()
	defer L.Close()

	for _, test := range tests {
		if err := L.DoFile(test); err != nil {
			return fmt.Sprintf("%s \n", err)
		}
	}

	return ""
}

func Demo() {
	L := lua.NewState()
	defer L.Close()

	// Load benchmark file
	if err := L.DoFile("benchmark.lua"); err != nil {
		fmt.Printf("%s \n", err)
		return
	}

	// Lua function in Go
	L.CallByParam(lua.P{
		Fn: L.GetGlobal("benchmark"),
		NRet: 1,
		Protect: true,
	}, lua.LNumber(10))
	fmt.Printf("print `depth` from Go: %v \n",  L.Get(-1))

	// Register Go function in Lua
	L.Register("count", func(l *lua.LState) int {
		L.Push(lua.LNumber(8))
		fmt.Println("echo `count` in Go")
		return 1
	})

	// Run Go function in Lua
	// L.Call(0, 1)
	L.DoString("countFromGo()")
}
