package design02

import "fmt"

func AbstractFactoryDesign() {
	woodenFactory, err := getFactory("wooden")
	if err != nil {
		fmt.Println(err)
	}
	chair := woodenFactory.createChair("greenPly")
	fmt.Println(chair.description())

}

func getFactory(factoryType string) (furnitureFactory, error) {
	switch factoryType {
	case "wooden":
		return &woodenFurnitureFactory{}, nil
	case "steel":
		return &steelFurnitureFactory{}, nil
	default:
		return nil, fmt.Errorf("no such factory")
	}
}

type furnitureFactory interface {
	createTable(company string) furniture
	createChair(company string) furniture
}

type woodenFurnitureFactory struct{}

func (wf *woodenFurnitureFactory) createTable(company string) furniture {
	return &furnitureType{
		name:     "wooden-table",
		company:  company,
		material: "wooden",
	}
}

func (wf *woodenFurnitureFactory) createChair(company string) furniture {
	return &furnitureType{
		name:     "wooden-chair",
		company:  company,
		material: "wooden",
	}
}

type steelFurnitureFactory struct{}

func (st *steelFurnitureFactory) createTable(company string) furniture {
	return &furnitureType{
		name:     "steel-table",
		company:  company,
		material: "steel",
	}
}

func (wf *steelFurnitureFactory) createChair(company string) furniture {
	return &furnitureType{
		name:     "steel-chair",
		company:  company,
		material: "steel",
	}
}

type furniture interface {
	description() string
}

type furnitureType struct {
	name     string
	company  string
	material string
}

func (ft *furnitureType) description() string {
	return fmt.Sprintf("%s made up of %s by %s", ft.name, ft.company, ft.material)
}
