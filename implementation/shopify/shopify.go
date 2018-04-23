package shopify

import (
	"github.com/Shopify/go-lua"
	"fmt"
)

func Run() string {
	L := lua.NewState()
	lua.OpenLibraries(L)
	if err := lua.DoFile(L, "benchmark.lua"); err != nil {
		return fmt.Sprintf("%s \n", err)
	}

	return ""
}
