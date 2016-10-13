package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/op/go-logging"
	"github.com/satori/go.uuid"
	"reflect"
	"strconv"
)

var myLogger = logging.MustGetLogger("diamond")

type DiamondChainCode struct {
}

type Owner struct {
	Id 			int32
	Name		string
	Address		string
	Company		string
}

type Diamond struct {
	Id			int32
	Owner_Id	int32
	Auth_Id		int32
	Weight		float32
	Height		float32
	Width		float32
}

type Authenticator struct {
	Id 			int32
	Name		string
	Company		string
}

type TradingPlatform struct {
	Id 			int32
	Name		string
	Website		string
}

type DiamondsOnTradePlat struct {
	Diamond_Id	int32
	Platform_Id int32
}

func (t *DiamondChainCode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

func (t *DiamondChainCode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

func (t *DiamondChainCode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

func main() {

}