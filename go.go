package main



import (
    "fmt"
    "time"
    
)







type BadChaincode struct {
    globalValue string
}
var globalValue = ""
var maps map[string]bool

func (t BadChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	go writeToLedger("data1")
	go writeToLedger("data2")
	
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
	

    fn, args := stub.GetFunctionAndParameters()
    if fn == "setValue" {
        t.globalValue = args[0]
		
        shim.PutState("key", []byte(t.globalValue))
		
		shim.PutState("key2", []byte(globalValue))
				
        return shim.Success([]byte("success"))
    } else if fn == "getValue" {
        value, _ := shim.GetState("key")
        return shim.Success(value)
    }
	
	
	for i := range maps{
		var aa=maps[i]
	}
	
	
	
	
	
    return shim.Error("not a valid function")
}


//no concurrency

func writeToLedger(stub shim.ChaincodeStubInterface, data string) {
	stub.PutState("key", []byte(data))
}



func main() {
//blank
return

}





