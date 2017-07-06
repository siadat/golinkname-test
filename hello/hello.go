package hello

import _ "unsafe" // for go:linkname

//go:linkname hola github.com/siadat/golinkname-test/greet.hellofunc
func hola() string {
	return "Hello from hola() in hello package"
}

//go:linkname structPrivateMethod github.com/siadat/golinkname-test/greet.A.privateMethod
func structPrivateMethod() string {
	return "Hello from structPrivateMethod() in hello package"
}
