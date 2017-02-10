#!/usr/bin/env bash
set -e

function go-tool-compile-pkg() {
  local importpath=$1
  local pkg=$(basename $importpath)
  local out=$GOPATH/src/$importpath.a
  go tool compile -I $GOPATH/src -pack -o $out $GOPATH/src/$importpath/*.go
}

function go-tool-link-pkg() {
  local importpath=$1
  local pkg=$(basename $importpath)
  local out=$GOPATH/src/$importpath/$pkg
  go tool link -o $out -L $GOPATH/src $GOPATH/src/$importpath.a
}

go-tool-compile-pkg github.com/siadat/golinkname-test/hello
go-tool-compile-pkg github.com/siadat/golinkname-test/greet
go-tool-compile-pkg github.com/siadat/golinkname-test
go-tool-link-pkg github.com/siadat/golinkname-test
