#### issue-1

while using
`fmt.Scanln(),fmt.Scan()`
it gives unexpected behavior when user have spaces in input or user press backspace
while giving input to the application through the cli
below is the example

```go
	fmt.Println("enter the street address where conference is held : ")
	fmt.Scanln(&streetAddress)
	fmt.Println("enter the pin code where conference is held: ")
	fmt.Scanln(&pinCode)
```

or we can use

```go
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter the street address where the conference is held: ")
    streetAddress, _ := reader.ReadString('\n')
    fmt.Println("Enter the pin code where the conference is held: ")
    pinCode, _ := reader.ReadString('\n')
    // Trim the newline character from the inputs
    streetAddress = streetAddress[:len(streetAddress)-1]
    // trim /r from the user input in case the os is windows
    pinCode = strings.TrimSpace(pinCode[:len(pinCode)-1])
```

> note while taking input from the user in widows. the input ends with "\r\n", where \r is the return carriage escape sequence and \n is the new line escape sequence character. so to remove both the escape sequence we have use `strings.TrimSpace()` to remove \r and `inputVal[:len(inputVal)-1]`
