package main

import "fmt"

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
