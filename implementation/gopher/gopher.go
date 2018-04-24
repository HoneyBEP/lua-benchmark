package gopher

import (
	"fmt"
	"github.com/yuin/gopher-lua"
)

/**

Gopher uses Lua v5.1

 */

func Run() string {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile("lua-tests/v5.1/all.lua"); err != nil {
		return fmt.Sprintf("%s \n", err)
	}

	return ""
}
