package main

import "fmt"

type example struct {
	exampleValue *exampleValue
}

type exampleValue struct {
	value int
}

func makeExampleSlice() []example {
	return []example{
		{exampleValue: &exampleValue{value: 0}},
		{exampleValue: &exampleValue{value: 0}},
	}
}

func output(name string, exampleSlice []example) {
	fmt.Printf("%s: %+v\n", name, exampleSlice)
}

func main() {
	examples := makeExampleSlice()

	exampleSlice1 := append(examples, example{exampleValue: &exampleValue{value: 1}})
	exampleSlice2 := append(examples, example{exampleValue: &exampleValue{value: 2}})

	output("example1", exampleSlice1)
	output("example2", exampleSlice2)
}
