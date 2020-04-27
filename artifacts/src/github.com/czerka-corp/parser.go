package main

import (
	"fmt"
	"encoding/json"
)

func inputSanitation(args []string, err error) {
	err = SanitizeArguments(args)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func ParseMercenary(args []string, err error) (Mercenary) {
	inputSanitation(args, err)

	var mercenary Mercenary
	mercenary.Id = args[0]
	mercenary.FirstName = args[1]
	mercenary.LastName = args[2]
	mercenary.Race = args[3]
	mercenary.Class = args[4]

	fmt.Println("Mercenary: ", mercenary)
	return mercenary
}

func ParseOffice(args []string, err error) (Office) {
	inputSanitation(args, err)
	planets := args[1]
	mercenaries := args[3]

	planetStruct := Planet{}
	planetStructAsBytes := []byte(planets)
	err = json.Unmarshal(planetStructAsBytes, &planetStruct)
	if err != nil {
		fmt.Println("unmarshalling Planet failed=", err)
	}

	mercenaryStruct := Mercenary{}
	mercenaryStructAsBytes := []byte(mercenaries)
	err = json.Unmarshal(mercenaryStructAsBytes, &mercenaryStruct)
	if err != nil {
		fmt.Println("unmarshalling Mercenary failed=", err)
	}

	var office Office
	office.Id = args[0]
	office.Planets = append(office.Planets, planetStruct)
	office.Location = args[2]
	office.Mercenaries = append(office.Mercenaries, mercenaryStruct)
	office.Head = args[4]

	fmt.Println("Office: ", office)
	return office
}
