package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AdressesList: []Adresses{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in adresses
	adressesIdMap := make(map[uint64]bool)
	adressesCount := gs.GetAdressesCount()
	for _, elem := range gs.AdressesList {
		if _, ok := adressesIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for adresses")
		}
		if elem.Id >= adressesCount {
			return fmt.Errorf("adresses id should be lower or equal than the last id")
		}
		adressesIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
