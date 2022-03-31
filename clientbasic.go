package main

import(
	"ethos/altEthos"
	"ethos/syscall"
	"log"
)

func init() {
	SetupMyRpcIncrementReply (incrementReply)
}

func incrementReply (count uint64) (MyRpcProcedure ) {
	log.Printf( "myRpcClient : Received Increment Reply : %v\n" ,count)
	return nil
} 


func main () {

	altEthos.LogToDirectory ("test/myRpcClient")

	log.Printf("myRpcClient : before call \n")

	fd , status := altEthos.IpcRepeat("myRpc" , "" , nil)
	if status!= syscall.StatusOk {
		log.Printf( " Ipc failed : %v\n" , status)
		altEthos.Exit(status)
	}

	call := MyRpcIncrement {}
	status = altEthos.ClientCall(fd , &call)
	if status!= syscall.StatusOk {
		log.Printf( "clientCallfailed : %v\n" , status)
		altEthos.Exit(status)
	}

	log.Printf( "myRpcClient : done\n" )
}
