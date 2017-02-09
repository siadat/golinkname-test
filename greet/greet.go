package greet

import _ "github.com/siadat/golinkname-test/hello"

func hellofunc() string // provided by hello package

func Greet() string {
	return hellofunc()
}
