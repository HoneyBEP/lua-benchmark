package shopify

import (
	"fmt"
	"github.com/Shopify/go-lua"
)

func Run(tests []string) string {
	L := lua.NewState()
	lua.OpenLibraries(L)

	for _, test := range tests {
		if err := lua.DoFile(L, test); err != nil {
			return fmt.Sprintf("%s \n", err)
		}
	}

	return ""
}
