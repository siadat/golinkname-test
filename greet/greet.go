package greet

import _ "github.com/siadat/golinkname-test/hello"

func hellofunc() string // provided by hello package

func Greet() string {
	return hellofunc()
}

type A struct{}

func (a A) privateMethod() string // provided by hello package

func (a A) ExportedMethod() string {
	return a.privateMethod()
}
