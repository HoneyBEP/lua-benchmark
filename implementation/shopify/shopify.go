package shopify

import (
	"github.com/Shopify/go-lua"
	"fmt"
)

/**

Shopify uses Lua 5.2

 */

func Run() string {
	L := lua.NewState()
	lua.OpenLibraries(L)
	if err := lua.DoFile(L, "lua-tests/v5.2/all.lua"); err != nil {
		return fmt.Sprintf("%s \n", err)
	}

	return ""
}
