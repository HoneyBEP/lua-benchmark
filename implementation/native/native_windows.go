// +build windows

package native

import (
	"fmt"
)

func Run(_  []string) string {
	return fmt.Sprintf("Unavailable")
}

func Demo() {
	fmt.Printf("Unavailable\n")
}