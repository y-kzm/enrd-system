package meas_client

/*
#cgo LDFLAGS: -L../igi-ptr -lptr-client
#include <../igi-ptr/setsignal.h>
#include <../igi-ptr/client_includes.h>
*/
import (
	"C"
)

func EstimateClient(packet_num int, probe_num int, packet_size int, src_addr string, dst_addr string) float64 {
	res := C.main_client(C.int(packet_num), C.int(probe_num), C.int(packet_size), C.CString(src_addr), C.CString(dst_addr))

	return float64(res)
}
