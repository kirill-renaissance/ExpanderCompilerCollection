package go_libs

/*
#cgo LDFLAGS: ${SRCDIR}/lib/libec_go_lib.a -ldl
#include <stdlib.h>
#include "./lib/ec_go.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

func CompileWithRustLib(s []byte, configId uint64) ([]byte, []byte, error) {
	in := C.ByteArray{data: (*C.uint8_t)(C.CBytes(s)), length: C.uint64_t(len(s))}
	defer C.free(unsafe.Pointer(in.data))

	cr := C.compile(in, C.uint64_t(configId))

	defer C.free(unsafe.Pointer(cr.ir_witness_gen.data))
	defer C.free(unsafe.Pointer(cr.layered.data))
	defer C.free(unsafe.Pointer(cr.error.data))

	irWitnessGen := C.GoBytes(unsafe.Pointer(cr.ir_witness_gen.data), C.int(cr.ir_witness_gen.length))
	layered := C.GoBytes(unsafe.Pointer(cr.layered.data), C.int(cr.layered.length))
	errMsg := C.GoBytes(unsafe.Pointer(cr.error.data), C.int(cr.error.length))

	if len(errMsg) > 0 {
		return nil, nil, errors.New(string(errMsg))
	}

	return irWitnessGen, layered, nil
}
