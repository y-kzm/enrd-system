package main

/*
#cgo LDFLAGS: -L../igi-ptr -lptr-server
#include <../igi-ptr/setsignal.h>
#include <../igi-ptr/server_includes.h>
*/
import "C"

func EstimateServer() {
	//C.Usage()
	C.main_server()
}

func main() {
	EstimateServer()
}
