package design01

import "fmt"

func FactoryDesign() {
	fmt.Println("an example of factory design pattern")
	wff := woodenFurnitureFactory{}
	var woodenChair furnitureType = wff.createChair()
	fmt.Println(woodenChair.description())
	var woodenTable furnitureType = wff.createTable()
	fmt.Println(woodenTable.description())
}

type furnitureFactory interface {
	createChair() furniture
	createTable() furniture
}

type woodenFurnitureFactory struct{}

func (wff *woodenFurnitureFactory) createChair() furnitureType {
	return &furniture{
		item:     "Chair",
		material: "wooden",
	}
}

func (wff *woodenFurnitureFactory) createTable() furnitureType {
	return &furniture{
		item:     "table",
		material: "wooden",
	}
}

type furnitureType interface {
	description() string
}

type furniture struct {
	item     string
	material string
}

func (ch *furniture) description() string {
	return fmt.Sprintf("a %v made of %v \n", ch.item, ch.material)
}
