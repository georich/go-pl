// Better implentation of echo
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))

	// Exercise 1.1 Modify to also print os.Args[0]
	fmt.Println(strings.Join(os.Args[0:], " "))
}
