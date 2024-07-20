package main


import (
	"fmt"

	"github.com/monksmojo/go-dojo/go-design-pattern/design01"
	"github.com/monksmojo/go-dojo/go-design-pattern/design02"
	"github.com/monksmojo/go-dojo/go-design-pattern/design03"
	"github.com/monksmojo/go-dojo/go-design-pattern/design04"
)

func main() {
	fmt.Println("welcome to the design pattern")
	design01.FactoryDesign()
	design02.AbstractFactoryDesign()
	design03.BuilderDesign()
	design04.AdapterDesign()
}
