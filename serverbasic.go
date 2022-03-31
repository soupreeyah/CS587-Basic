package main

import (
	"ethos/syscall"
	"ethos/altEthos"
	"log"
)

var myRpc_increment_counter uint64 = 0

func init() {
	SetupMyRpcIncrement (increment)
}

func increment () (MyRpcProcedure) {
	log.Printf( "myRpcService : called increment \n" )
	myRpc_increment_counter++
	return &MyRpcIncrementReply {myRpc_increment_counter}
}

func main () {

	altEthos.LogToDirectory("test/myRpcService")

	listeningFd, status:= altEthos.Advertise("myRpc")
	if status!= syscall.StatusOk {
		log.Printf( "Advertising service failed : %s \n" , status)
		altEthos.Exit(status)
	}

	for {
		_, fd , status := altEthos.Import(listeningFd)
		if status!=syscall.StatusOk {
			log.Printf( "Error calling Import : %v\n" , status)
			altEthos.Exit(status)
		}	

		log.Printf("myRpcService : new connection accepted \n")

		t := MyRpc{}
		altEthos.Handle (fd,&t)
	}

}
