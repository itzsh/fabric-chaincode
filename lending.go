/*
This is a chaincode that used to lending
*/

package main

import {
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
}

type Lending struct {
}

//Init function
func (t *Lending) Init(stub shim.ChaincodeStubInterface) ([]byte, error) {
	_, args := stub.GetFunctionAndParameters()

	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	//Create asset table
	err := stub.CreateTable("AssetsTable", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "assetID", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "assetOwner", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "assetStatus", Type: shim.ColumnDefinition_INT32, Key: true},
		&shim.ColumnDefinition{Name: "assetDataInfo", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "assetAlterDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "assetKeyMark", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "assetKeyMarkValue", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, fmt.Errorf("Failed creating AssetsTable table, [%v]", err)
	}
	return nil, nil
}

func (t *Lending) Invoke(stub shim.ChaincodeStubInterface) ([]byte, error) {
	function, args := stub.GetFunctionAndParameters()
	if function == "insertAsset" {
		return t.insertAsset(stub, args)
	} else if function == "queryAsset" {
		return t.queryAsset(stub, args)
	}

	return nil, errors.New("Invalid invoke function name. Expecting \"insertAsset\" \"queryAsset\"")
}

func (t *Lending) insertAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 7 {
		return nil, errors.New("Incorrect number of arguments. Expecting 7")
	}

	assetID := args[0]
	assetOwner := args[1]
	assetStatus := args[2]
	assetDataInfo := args[3]
	assetAlterDate := args[4]
	assetKeyMark := args[5]
	assetKeyMarkValue := args[6]

	//是否要加上权限管理，只有部署链码的人才能调用
	ok, err = stub.InsertRow ("AssetsTable", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: assetID}},
			&shim.Column{Value: &shim.Column_String_{String_: assetOwner}},
			&shim.Column{Value: &shim.Column_Int32{Int32: assetStatus}}},
			&shim.Column{Value: &shim.Column_String_{String_: assetDataInfo}},
			&shim.Column{Value: &shim.Column_String_{String_: assetAlterDate}},
			&shim.Column{Value: &shim.Column_String_{String_: assetKeyMark}},
			&shim.Column{Value: &shim.Column_String_{String_: assetKeyMarkValue}},
	})

	if !ok && err == nil {
		return nil, errors.New("Asset was already insert.")
	}

	return nil, err
}

func (t *Lending) queryAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	assetID := args[0]

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: assetID}}
	columns = append(columns, col1)

	row, err := stub.GetRow("AssetsTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed retrieving asset [%s]: [%s]", asset, err)
	}

	assetOwner := row.Columns[1].GetString_()
	fmt.Printf("The assetOwner is %s", assetOwner)

	return nil, nil
}

func main() {
	err := shim.Start(new(Lending))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
