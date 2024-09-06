### in this topic we will learn about the

#### main package

- packages in go are the directories that contain multiple go files having similar function,variable and constants

- the primary and mandatory package in go is `package main`
  which contains `func main(){}` function
- `func main(){}` function act as the entry point to the go compiler for the execution of the program for example

```go
  package main
  import "fmt"
  func main() {
    fmt.Println("Yo Hello World!")
    }

```
- `func` is keyword used to create a function followed function name and function definition 
- `fmt` package contains multiple function the one which we are using are used for printing output to the console

### how to write unit test in go (approach-1)
- It is good to put the code/logic to test in separate functions. 
- I have learned following steps while writing a test in go.
  - creating a `name_test.go` file where we write our unit test.
  - the unit test function start with `TestNameOfFunction(t *testing.T)` taking 
  `t *testing.T` as the only argument.
  - we are using `import "testing"` package. 


