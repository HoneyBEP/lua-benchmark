// +build !windows

package native

import (
	"fmt"
	"github.com/aarzilli/golua/lua"
)

func Run(tests []string) string {
	L := lua.NewState()
	defer L.Close()
	L.OpenLibs()

	for _, test := range tests {
		if err := L.DoFile(test); err != nil {
			return fmt.Sprintf("%s \n", err)
		}
	}
	return ""
}

func Demo() {
	fmt.Printf("Not implemented yet")
}