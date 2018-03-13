package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
	"fmt"
)

func putMercenary(mercenary Mercenary, stub shim.ChaincodeStubInterface, err error) (pb.Response) {
	mercenaryAsBytes, _ := json.Marshal(mercenary)
	err = stub.PutState(mercenary.Id, mercenaryAsBytes)
	if err != nil {
		fmt.Println("Could not store mercenary")
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func InsertMercenary(stub shim.ChaincodeStubInterface, args []string) (pb.Response) {
	var err error
	fmt.Println("starting InsertMercenary")

	if len(args) != 12 {
		return shim.Error("InsertMercenary(): Incorrect number of arguments. Expecting 12")
	}

	mercenary := ParseMercenary(args, err)
	_, err = GetMercenary(stub, mercenary.Id)

	if err == nil {
		return shim.Error("This mercenary already exists - " + mercenary.Id)
	}

	putMercenary(mercenary, stub, err)
	fmt.Println("- end InsertMercenary")
	return shim.Success(nil)
}

func UpdateMercenary(stub shim.ChaincodeStubInterface, args []string) (pb.Response) {
	var err error
	fmt.Println("starting UpdateMercenary")

	if len(args) != 12 {
		return shim.Error("UpdateMercenary(): Incorrect number of arguments. Expecting 12")
	}

	mercenary := ParseMercenary(args, err)
	_, err = GetMercenary(stub, mercenary.Id)

	if err != nil {
		return shim.Error("Failed to get mercenary - " + mercenary.Id)
	}

	putMercenary(mercenary, stub, err)
	fmt.Println("- end UpdateMercenary")
	return shim.Success(nil)
}

func putOffice(office Office, stub shim.ChaincodeStubInterface, err error) (pb.Response) {
	officeAsBytes, _ := json.Marshal(office)
	err = stub.PutState(office.Id, officeAsBytes)
	if err != nil {
		fmt.Println("Could not store office")
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func InsertOffice(stub shim.ChaincodeStubInterface, args []string) (pb.Response) {
	var err error
	fmt.Println("starting InsertOffice")

	if len(args) != 5 {
		return shim.Error("InsertOffice(): Incorrect number of arguments. Expecting 5")
	}

	office := ParseOffice(args, err)
	_, err = GetOffice(stub, office.Id)

	if err == nil {
		return shim.Error("This office already exists - " + office.Id)
	}

	putOffice(office, stub, err)
	fmt.Println("- end InsertOffice")
	return shim.Success(nil)
}

func UpdateOffice(stub shim.ChaincodeStubInterface, args []string) (pb.Response) {
	var err error
	fmt.Println("starting UpdateOffice")

	if len(args) != 5 {
		return shim.Error("UpdateOffice(): Incorrect number of arguments. Expecting 5")
	}

	office := ParseOffice(args, err)
	_, err = GetOffice(stub, office.Id)

	if err != nil {
		return shim.Error("Failed to get office - " + office.Id)
	}

	putOffice(office, stub, err)
	fmt.Println("- end UpdateOffice")
	return shim.Success(nil)
}
