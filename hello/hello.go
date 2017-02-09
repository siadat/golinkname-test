package hello

import _ "unsafe" // for go:linkname

//go:linkname hola github.com/siadat/golinkname-test/greet.hellofunc
func hola() string {
	return "Hello from hello package"
}
