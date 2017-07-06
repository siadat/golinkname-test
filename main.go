package main

import (
	"fmt"

	"github.com/siadat/golinkname-test/greet"
)

func main() {
	fmt.Println(greet.Greet())
	fmt.Println(greet.A{}.ExportedMethod())
}
