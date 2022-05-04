// This program prints its command-line arguments in an ordered fashion.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("%d | %s\n", i, arg)
	}
}
