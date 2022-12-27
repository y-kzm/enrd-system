package main

/*
#cgo LDFLAGS: -L../igi-ptr -lptr-client
#include <../igi-ptr/setsignal.h>
#include <../igi-ptr/client_includes.h>
*/
import "C"
import "fmt"

//func EstimateClient(packet_num int, probe_num int, packet_size int, src_addr string, dst_addr string) float64 {
func EstimateClient() {
	res := C.main_client()
	fmt.Printf("PTR: %.3f\n", float64(res))
}

func main() {
	EstimateClient()
	//fmt.Print(res)
}
