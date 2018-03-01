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