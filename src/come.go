package main

import (
	"os"
	"compiler"
	"vm"
	"io"
	"strings"
)

func main() {

	// load file
	var (
		filename string
		f io.Reader
		err error
	)
	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		filename = ""
	}
	f, err = os.Open(filename)
	if err != nil {
		f = strings.NewReader(compiler.TestSrcCode)
	}

	// compile
	co := compiler.Compile(f)

	// run
	vm.Run(co)
}
