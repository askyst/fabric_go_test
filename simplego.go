package main



import (
    "fmt"
    "time"
    
)







type BadChaincode struct {
    globalValue string
}


func (t BadChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {


	
	key := "key"
    data := "data"
    err := stub.PutState(key, []byte(data)) //here
    if err != nil {
        return shim.Error("could not write new data")
    }
    respone, err := stub.GetState(key)  //here
    if err != nil {
        return shim.Error("could not read data")
    }
	

}





func main() {
//blank
return

}





