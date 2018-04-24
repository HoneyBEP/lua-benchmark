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
