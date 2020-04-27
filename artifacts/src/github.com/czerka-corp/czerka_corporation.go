package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"fmt"
	"strconv"
)

type CzerkaContract struct {
}

type Mercenary struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Race      string `json:"race"`
	Class     string `json:"class"`
}

type Office struct {
	Id          string `json:"id"`
	Location    string `json:"location"`
	Mercenaries []Mercenary `json:"mercenaries"`
	Supplies 	[]Supply `json: "supplies"`
	Head        string `json:"head"`
}

type Planet struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Offices 	[]Office `json: "offices"`
	Region      string `json:"region"`
	Sector      string `json:"sector"`
	System      string `json:"system"`
	Coordinates string `json:"coordinates"`
}

type Supply struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Status string `json:"status"`
}

func main() {
	err := shim.Start(new(CzerkaContract))
	if err != nil {
		fmt.Printf("Error starting chaincode - %s", err)
	}
}

//
//Init
//
func (t *CzerkaContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Init Is Starting Up")
	funcName, args := stub.GetFunctionAndParameters()
	var number int
	var err error
	txId := stub.GetTxID()

	fmt.Println("Transaction ID:", txId)
	fmt.Println("  GetFunctionAndParameters() function:", funcName)
	fmt.Println("  GetFunctionAndParameters() args count:", len(args))
	fmt.Println("  GetFunctionAndParameters() args found:", args)

	if len(args) == 1 {
		fmt.Println("  GetFunctionAndParameters() arg[0] length", len(args[0]))

		// expecting arg[0] to be length 0 for upgrade
		if len(args[0]) == 0 {
			fmt.Println("  Args[0] is empty...")
		} else {
			fmt.Println("  Args[0] is not empty")

			// convert numeric string to integer
			number, err = strconv.Atoi(args[0])
			if err != nil {
				return shim.Error("Expecting a numeric string argument to Init() for instantiate")
			}

			err = stub.PutState("test", []byte(strconv.Itoa(number)))
			if err != nil {
				return shim.Error(err.Error())
			}
		}
	}

	// showing the alternative argument shim function
	alt := stub.GetStringArgs()
	fmt.Println("  GetStringArgs() args count:", len(alt))
	fmt.Println("  GetStringArgs() args found:", alt)

	fmt.Println("Ready for action") //self-test pass
	return shim.Success(nil)
}

func (t *CzerkaContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println(" ")
	fmt.Println("starting invoke, for - " + function)

	if function == "init" {
		//initialize the chaincode state, used as reset
		return t.Init(stub)
	} else if function == "InsertMercenary" {
		return InsertMercenary(stub, args)
	} else if function == "UpdateMercenary" {
		return UpdateMercenary(stub, args)
	} else if function == "InsertOffice" {
		return InsertOffice(stub, args)
	} else if function == "UpdateOffice" {
		return UpdateMercenary(stub, args)
	}
	// error out
	fmt.Println("Received unknown invoke function name - " + function)
	return shim.Error("Received unknown invoke function name - '" + function + "'")
}
