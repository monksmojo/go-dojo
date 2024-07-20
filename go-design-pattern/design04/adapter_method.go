package design04

import "fmt"

func AdapterDesign() {
	airportUnit := unit{
		location: "airport",
	}
	tabletDevice := &tablet{
		name:         "samsung tablet",
		chargingType: " type c",
	}
	phoneDevice := &phone{
		name:         "nokia2241",
		chargingType: "type b",
	}
	phoneAdapterExtension := &phoneAdapter{
		ph: phoneDevice,
	}
	fmt.Println(airportUnit.chargerTypeC(tabletDevice))
	fmt.Println(airportUnit.chargerTypeC(phoneAdapterExtension))
}

// target interface that client can interact with
type chargingStation interface {
	chargerDevice() string
}

// concrete class that interacts with the
type tablet struct {
	name         string
	chargingType string
}

func (t *tablet) chargerDevice() string {
	return fmt.Sprintf("charging %v at the type c dock", t.name)
}

type phoneAdapter struct {
	ph *phone
}

func (pa *phoneAdapter) chargerDevice() string {
	return pa.ph.chargeTypeB()
}

// legacy adoptee
type phone struct {
	name         string
	chargingType string
}

func (p *phone) chargeTypeB() string {
	return fmt.Sprintf("charging %v at the type b dock", p.name)
}

// client
type unit struct {
	location string
}

func (u *unit) chargerTypeC(cs chargingStation) string {
	return cs.chargerDevice()
}
