package gopher

import (
	"github.com/yuin/gopher-lua"
	"fmt"
)

func Run() string {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile("benchmark.lua"); err != nil {
		return fmt.Sprintf("%s \n", err)
	}

	return ""
}
