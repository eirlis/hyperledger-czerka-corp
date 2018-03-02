package main

import (
	"strconv"
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
)

func SanitizeArguments(strs []string) error {
	for i, val := range strs {
		if len(val) <= 0 {
			return errors.New("Argument " + strconv.Itoa(i) + " must be a non-empty string")
		}
	}
	return nil
}

func GetMercenary(stub shim.ChaincodeStubInterface, id string) (Mercenary, error) {
	var mercenary Mercenary
	mercenaryAsBytes, err := stub.GetState(id)
	if err != nil {
		return mercenary, errors.New("Failed to find mercenary - " + id)
	}
	json.Unmarshal(mercenaryAsBytes, &mercenary)
	if mercenary.Id != id {
		return mercenary, errors.New("Mercenary does not exist - " + id)
	}
	return mercenary, nil
}

func GetOffice(stub shim.ChaincodeStubInterface, id string) (Office, error) {
	var office Office
	officeAsBytes, err := stub.GetState(id)
	if err != nil {
		return office, errors.New("Failed to find office - " + id)
	}
	json.Unmarshal(officeAsBytes, &office)
	if office.Id != id {
		return office, errors.New("Office does not exist - " + id)
	}
	return office, nil
}

func GetPlanet(stub shim.ChaincodeStubInterface, id string) (Planet, error) {
	var planet Planet
	planetAsBytes, err := stub.GetState(id)
	if err != nil {
		return planet, errors.New("Failed to find planet - " + id)
	}
	json.Unmarshal(planetAsBytes, &planet)
	if planet.Id != id {
		return planet, errors.New("Planet does not exist - " + id)
	}
	return planet, nil
}

func GetSupply(stub shim.ChaincodeStubInterface, id string) (Supply, error) {
var supply Supply
	supplyAsBytes, err := stub.GetState(id)
	if err != nil {
		return supply, errors.New("Failed to find supply - " + id)
	}
	json.Unmarshal(supplyAsBytes, &supply)
	if supply.Id != id {
		return supply, errors.New("Supply does not exist - " + id)
	}
	return supply, nil
}