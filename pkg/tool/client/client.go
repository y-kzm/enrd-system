package tool

/*
#cgo LDFLAGS: -L../igi-ptr -lptr-client
#include <../igi-ptr/setsignal.h>
#include <../igi-ptr/client_includes.h>
*/
import "C"

func EstimateClient() {
    C.main_client()
}
