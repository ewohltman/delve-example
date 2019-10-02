# `delve-example`

Delve: [https://github.com/go-delve/delve](https://github.com/go-delve/delve)

---

## Debugging main application

### Start Delve
```
dlv debug delve-example.go
```

### Set breakpoints
```
(dlv) break delve-example.go:30
Breakpoint 1 set at 0x4ac131 for main.main() ./delve-example.go:30
```

```
(dlv) break delve-example.go:21
Breakpoint 2 set at 0x4abeaf for main.output() ./delve-example.go:21
```

### Run until first breakpoint
```
(dlv) continue
> main.main() ./delve-example.go:30 (hits goroutine(1):1 total:1) (PC: 0x4ac131)
    25:		examples := makeExampleSlice()
    26:	
    27:		exampleSlice1 := append(examples, example{exampleValue: &exampleValue{value: 1}})
    28:		exampleSlice2 := append(examples, example{exampleValue: &exampleValue{value: 2}})
    29:	
=>  30:		output("example1", exampleSlice1)
    31:		output("example2", exampleSlice2)
    32:	}
```

#### List function local variables
```
(dlv) locals
examples = []main.example len: 2, cap: 2, [...]
exampleSlice2 = []main.example len: 3, cap: 4, [...]
exampleSlice1 = []main.example len: 3, cap: 4, [...]
```

#### Print value of local variable
```
(dlv) print exampleSlice1[2]
main.example {
	exampleValue: *main.exampleValue {value: 1},}
```

### Run until second breakpoint
```
(dlv) continue
> main.output() ./delve-example.go:21 (hits goroutine(1):1 total:1) (PC: 0x4abeaf)
    16:			{exampleValue: &exampleValue{value: 0}},
    17:		}
    18:	}
    19:	
    20:	func output(name string, exampleSlice []example) {
=>  21:		fmt.Printf("%s: %+v\n", name, exampleSlice)
    22:	}
    23:	
    24:	func main() {
    25:		examples := makeExampleSlice()
    26:
```

#### List function arguments
```
(dlv) args
name = "example1"
exampleSlice = []main.example len: 3, cap: 4, [...]
```

#### Print value of function argument
```
(dlv) print exampleSlice[2]
main.example {
	exampleValue: *main.exampleValue {value: 1},}
```

---

## Debugging tests

### Start Delve
```
dlv test ./internal/pkg/example/
```

### Set breakpoints
```
(dlv) break example_test.go:9
Breakpoint 1 set at 0x547e5f for github.com/ewohltman/delve-example/internal/pkg/example.TestNew() ./internal/pkg/example/example_test.go:9
```

```
(dlv) break example_test.go:29
Breakpoint 2 set at 0x5480a0 for github.com/ewohltman/delve-example/internal/pkg/example.TestNewSlice() ./internal/pkg/example/example_test.go:29
```

### See remaining examples under `Debugging main application`

---

## Exit Delve
```
(dlv) exit
```