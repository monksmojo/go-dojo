### GO PROGRAMMING LANGUAGE

> Following are the references, books and videos which are helping me along my go journey to become a gopher and help me build this go-topics repository.

> BIG THANKS TO EACH CREATORS AND DEVELOPERS OF FOR CREATING THESE AWESOME REFERENCES/COURSES/TUTORIALS.

- [Learn Go With Tests](https://quii.gitbook.io/learn-go-with-tests "Learn Go with tests")

- [Go by Example](https://gobyexample.com/hello-world "Go by Example")

- [learn-golang by boot.dev](https://www.boot.dev/learn/learn-golang "learn-golang by boot.dev")

- [Concurrency in Go by Jake Wright](https://youtu.be/LvgVSSpwND8?list=PL9KOhO7TtG6YazaOAKlU_elPlRDQ2ht3n "Concurrency in Go by Jake Wright youtube video tutorial")

- [Go programming language (Go Class) by Matt KØDVB](https://youtu.be/iDQAZEJK8lI?list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6 "Go programming language (Go Class) by Matt KØDVB youtube video tutorial")

#### Quotes from the Creators and Developers

> if our basic tools and language in which we design and code our program is also complicated, the language itself becomes the part of the problem rather than the part of the solution ~ Tony Hoare

> It's programming if "clever" is a compliment but it's software engineering if "clever" is an accusation. ~ A Software Developer

#### THE NEXT GEN LANGUAGE

- **The GO PROGRAMMING LANGUAGE**
- **(GOLANG)**
- It originated at Google by ken Ken Thompson, Rob Pike, and others
- Go is about language designed in the service of software engineering

#### Features Of Go Language

- It is Free and open source
- It Supports Statically typed
- It Contains Memory safety
- It Comes With Garbage collection
- It is structural typing
- It can follow CSP-style
- It Gives Python/JS like readability and usability
- Strong support for multi-core and networked systems
  Concurrency, and more.
- It Gives set of tools that we can utilize
- Use for system Apps, Web Apps and cloud architecture
- Golang is partial OOPS {contains structs, does not support classes,overloading}
- lexer does lot of work so no semi colon ; in code
- GO is case sensitive
- With introduction to go 1.22.0 we can range over an integer in for loop (note: range on integer start with 0 till n-1)

#### Misses In GO

- try catch is missing
- inheritance is missing

#### Go vs Java

- GoLang is Complied
- Java on the other hand is somewhere in between compiled and interpreted as it creates binary which is executed by jvm (java virtual machine)

#### go toolchain

- go toolchain is the set of tools provided by go programming language
- they are bundled with the go distribution
- the go toolchain has the following commands
  - `go build` compiles the packages and dependencies, creates a executable binary and stores
  - `go run` compiles the packages and dependencies and executes the executable while saving it at a temporary place. it deletes the executable after executing it
  - `go install` Compiles and installs the packages, placing the binaries in $GOPATH/bin or $GOBIN.
  - `go get` Downloads and installs packages and dependencies from remote repositories.
  - `go mod` Manages modules and dependencies. It includes various subcommands like init, tidy, and download.
    - Example: `go mod init` creates a `go.mod` file to manage the dependencies.
    - Example: `go mod tidy` cleans up unused dependencies in the `go.mod` file.
  - `go clean` Removes object files and cached files created during the build process.
  - `go env` Displays or sets environment variables used by the Go toolchain.
  - `go list` Lists the packages, modules, or other information about the project.

#### Difference between `go install` vs `go get`

- `go install`: 
  - Use `go install` with a specific version to install executables under $GOPATH/bin or $GOBIN without affecting your project's dependencies in `go.mod` file.
  - compiles, install binaries and install packages from the packages
  - does not modify `go.mod` or `go.sum`
- `go get`: 
  - Use `go get` to add or update dependencies within your Go modules.
  - add or update your dependencies in go.mod file 
  - it affects your `go.mod` file.
  



#### To Run Go File We write

- go run <file_name>.go

#### Explaining the command go run main.go

- go run first compiles the Go source file(s) into machine code.
- It uses the Go compiler (gccgo) to perform this compilation.
- After compiling, it immediately executes the resulting binary.
- This execution happens in memory, not as a separate step after compilation.
- When the program finishes running, the compiled binary is automatically deleted.

#### Explaining the command go build -0 myApp main.go

- go build parses and compiles the source code into intermediate object files.
- the object files are linked to produce an executable binary.
- go build saves the resulting executable in the current directory (default) or a specified output path.


#### difference in go build and go run
- Aspect:	`go run`
- Primary: Purpose	Compile and execute immediately	
- Output:	Temporary executable (deleted after run)	
- Use Case:	Quick testing, running scripts	
- Performance:	Slightly slower due to cleanup steps	
- Flags Availability:	Limited flags specific to execution	

- Aspect:	`go build`
- Primary: Purpose Compile and produce a binary
- Output: Persistent executable file
- Use Case: Building applications for deployment
- Performance: Faster for repeated executions
- Flags: More extensive flags for building options
- @set GO_ENABLED=0 GOARCH=amd64 GOOS=linux & go build -o "./$(BUILD_DIR)/$(BINARY_NAME)-linux" "./cmd/app/main.go"

#### Go Has A Built In Documentation

- use command go help to see various tools available in go
- and to read about a particular go module we can write
- `go help <command_name>`

#### To See Go Install Path

- go env GOPATH

#### what is GOPATH

- `GOPATH` is an environment variable that specifies the root of your workspace. It is where Go looks for your source code, dependencies, and binaries.
- in linux machines it is mostly : /usr/local/go/bin
- A typical GOPATH workspace contains three directories:
  src: For source code.
  pkg: For compiled packages.
  bin: For compiled binaries.
- Before Go modules were introduced, GOPATH was the default way to manage Go projects.
- All Go code had to reside inside the GOPATH/src directory. Dependencies were manually managed or through tools like dep, and there was no native versioning of dependencies.
- Projects had to be kept inside the GOPATH workspace, leading to potential conflicts, especially when working on multiple projects with different versions of the same dependency.
- It didn't support versioned dependencies well, leading to challenges in managing dependencies across projects.

#### Introduction to `go.mod` file

#### Explaining why go module were needed

- go after the release of go 1.16 release go module
- go module help us to execute go program outside the `GoPath`
- The GOPATH environment variable in Go defines the root of your workspace. It is used by Go tools to organize your code and dependencies.

- It's important to note that while GOPATH is crucial for organizing your workspace and managing dependencies, it is not being removed. Instead, the Go ecosystem is transitioning from GOPATH development mode to Go modules, which offer a more robust and reproducible way to manage dependencies.

- modules are the way of storing internal and external libraries (3rd party)

- A module is defined by a go.mod file, which specifies the module's path, version, and its dependencies.

- the standard way of creating a go module is

```
go mod init github.com/<userName>/<projectName>
```

- Go modules are decentralized by nature. Unlike some other package management systems, there is no central repository for Go modules. Instead, modules can be hosted on any version control system, such as Git repositories hosted on services like GitHub, GitLab, or even on a local file system.

### packages vs modules in go

- modules are the way of storing internal and external libraries (3rd party)

- packages in go are the directories that contain multiple go files having similar function,variable and constants

- the primary and mandatory package in go is `package main`
  which contains `func main(){}` function
- `func main(){}` function act as the entry point to the go compiler for the execution of the program

  ```go
  package main
    import (
    "fmt"
    )
    func main() {
    fmt.Println("welcome to main driver class")
    }

  ```

- it is go practice to make a treat a single package as a single directory.
- custom package and directory name should be same
- for ex

```code
$ tree /03package
/03package
├── customPackage
│   |-- customMath.go // is declared in package customPackage
│   |-- customString.go // is declared in package customPackage
|
|--go.mod // which is the module file of /03package(also the package name) directory
|--main.go // which comes under package main
|
```

- to import the functions and variables from package customPackage into package main
- we have to import in the following manner

  - "module_name/package_name"
  - for example "03package/customPackage"

- only the public methods and variables can be shared between the packages
- in go we declare public methods and variables with _TitleCasing_ ex func GetValue(){}
- function and variables declared in multiple go file but inside same package can be shared without import

### Introduction to `go.work` file

#### Working With Multi Module Workspace

- when working with projects with multiple modules go.mod files
- gopl needs a central file to manage all your sub module
- For that we use `go work tool`
- From Go 1.18 onwards there is native support for multi-module workspaces. This is done by having a go.work file present in your parent directory.

-For a directory structure such as:

```code
$ tree /my/parent/dir
/my/parent/dir
├── project-one
│   ├── go.mod
│   ├── project-one.go
│   └── project-one_test.go
└── project-two
    ├── go.mod
    ├── project-two.go
    └── project-two_test.go
```

- Create and populate the file by executing go work:

```code
cd /my/parent/dir
go work init
go work use project-one
go work use project-two
```

- This will add a go.work file in your parent directory that contains a list of directories you marked for usage:

### how to write unit test in go (approach-1)

- It is good to put the code/logic to test in separate functions.
- I have learned following steps while writing a test in go.

  - creating a `name_test.go` file where we write our unit test.
  - the unit test function start with `TestNameOfFunction(t *testing.T)` taking
    `t *testing.T` as the only argument.
  - we are using `import "testing"` package.

- It may change as we advance

### what is lexer

- helps you to follow the grammar of the language
- checks syntax
- works pre compilation to check if you are following the rules

### Types and Variables in go

- Types are case sensitive
- Types helps us to defile the public and private accessibility of the function or variable
- for example
  ` fmt.Println("Yo Hello World!")`

- here 'P' is in upperCase defining that Println is a public function

- Variables of a type should be declared in advance before assigning -- giving the go the feature of type safe langue

- also go help us to declare type of variable on the run.

- wwe can say almost everything in golang is type

### Rules of naming Variables

- A variable name consists of alphabets, digits, and an underscore.
- Variables cannot have other symbols ( $, @, #, etc).
- Variable names cannot begin with a number.
- A variable name cannot be a reserved word as they are part of the Go syntax like int, type, for, etc.

### Constant in Go

- Constants are the fixed values that cannot be changed once declared. In Go, we use the const keyword to create constant variables. For example,

### Types GoLang Support

#### Basic Types

- String
- Bool
- Integer
- Floating
- Complex

#### Advance Types

- Array
- Slices
- Map
- Struts
- Pointers

### Why we say almost everything is type

- Functions and Channels are also Types

### Format specifiers

- Go language provides various format specifiers for different types of variables. Here are some commonly used format specifiers in Go:

- %d - decimal integer
- %f - floating-point number
- %s - string
- %t - boolean
- %v - any value
- %T - type of the variable
- %p - pointer

### In go it is easy to create an executable depending on the operating system

- type the `go env` command to the various environment variable supported by the go
- To create the environment in go
- for windows we can type `GOOS='windows' go build`
- for linux we can type `GOOS='linux' go build`
- for mac osx we can type `GOOS='darwin' go build`

### Golang memory management

- In golang memory management take place automatically
- But golang provide us 2 methods for memory management

1. new()
   - allocate memory to the variable but doest not initialize the variable with value
   - you will get a memory address which you can access using pointer
   - zero storage - no value will be initialized
2. make()
   - allocate memory to the variable and initialize the variable with default value
   - you will get a memory address which you can access using pointers
   - zero storage - no value will be initialized

- garbage collector will call automatically when call out of scope

- runtime is the package which tells the how go runtime works and how its garbage collector works
- there is a threshold value in go which triggers the go garbage collector when the threshold crosses

### Pointers

pointers are the variables which store memory location of the other variables

### why pointers were created

- when we pass variables as an argument to the function.
- the arguments are the copy of the variables passed to the function
- but sometimes we want to work/update the original variable in that case we use pointers
- pointers when passed in the function they pass the address of the variable.
- by default value of pointer in go is nil

### arrays in go language

- An array is a collection of elements of the same type.
- It is a fixed-size container that can hold a specific number of elements.
- Once an array is declared, its size cannot be changed.
- The elements of an array can be of any type such as int, float, bool, string, struct, etc.

### Go Print statement

- Three functions to print out the messages in the go
- fmt.Print()
- fmt.Println()
- fmt.Printf()

- all these packages are inside fmt package
