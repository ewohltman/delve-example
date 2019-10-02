package main

import (
	"fmt"

	"github.com/ewohltman/delve-example/internal/pkg/example"
)

func output(name string, exampleSlice []*example.Struct) {
	fmt.Printf("%s: %+v\n", name, exampleSlice)
}

func main() {
	examples := example.NewSlice()

	exampleStruct1 := example.New(1)
	exampleStruct2 := example.New(2)

	exampleSlice1 := append(examples, exampleStruct1)
	exampleSlice2 := append(examples, exampleStruct2)

	output("example1", exampleSlice1)
	output("example2", exampleSlice2)
}
