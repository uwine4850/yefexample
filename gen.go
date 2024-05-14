package main

import "github.com/uwine4850/yefgo/codegen"

func main() {
	err := codegen.Generate("pygen/pygen.yaml", "github.com/uwine4850/yefgotest")
	if err != nil {
		panic(err)
	}
}
