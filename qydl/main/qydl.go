package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"fmt"
	"encoding/json"
)

// struct
type qydlChaincode struct {

}

type qydl struct {
	ObjectType  string
	Number string
	Name string
	UpdateTime string
	UpdateSymbol string
	Owner string
}

func (t *qydlChaincode) Init (stub shim.ChaincodeStubInterface) peer.Response  {

	return shim.Success(nil)

}

func (t *qydlChaincode) Invoke (stub shim.ChaincodeStubInterface) peer.Response  {

	fn , args := stub.GetFunctionAndParameters()

	if fn == "addQydl"{
		return t.addQydl(stub,args)
	}else if fn =="readQydl"{
		return t.readQydl(stub,args)
	}

	return shim.Error("没有对应的方法!")

}

func (t *qydlChaincode) addQydl(stub shim.ChaincodeStubInterface, args [] string) peer.Response {
	qydlNumber := args[0]
	qydlAsBytes , err := stub.GetState(qydlNumber)

	if err!=nil {
		return shim.Error(err.Error())
	}
	if qydlAsBytes != nil{

		return shim.Error("number 已经存在")
	}

	name := args[1]
	time := args[2]
	symbol := args[3]
	owner := args[4]

	objectType := "qydl"

	qydl := &qydl{objectType,qydlNumber,name,time,symbol,owner}

	qydlJsonAsBytes , err := json.Marshal(qydl)




	err = stub.PutState(qydlNumber,qydlJsonAsBytes)

	if err!= nil {
		return shim.Error(err.Error())
	}

	fmt.Println("add===========")
	fmt.Println(qydlJsonAsBytes)
	fmt.Println("--------")
	fmt.Println(args[1])

	return shim.Success(nil )
}

func (t *qydlChaincode) readQydl(stub shim.ChaincodeStubInterface, args [] string) peer.Response  {
	qydlNumber := args[0]
	rstr , err := stub.GetState(qydlNumber)

	if err !=nil{
		return shim.Error(err.Error())
	} else if rstr == nil{
		return shim.Error("信息不存在!")
	}

	fmt.Println("==read============")
	fmt.Println(rstr)

	return shim.Success(rstr)

	
}


// del
func (t *qydlChaincode) delQydl(stub shim.ChaincodeStubInterface,args[] string) peer.Response  {

	return shim.Success(nil)
}



func main()  {
	 err := shim.Start(new(qydlChaincode))

	 if err!=nil{
	 	shim.Error(err.Error())
	 	fmt.Println(err.Error())
	 	fmt.Println("qydlChaincode start error1")
	 }
}