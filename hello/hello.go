package hello

// Importing unsafe just because compiler says: go:linkname only allowed in Go files that import "unsafe"
import _ "unsafe"

//go:linkname hello greet.hello
func hello() string {
	return "HELLO"
}
