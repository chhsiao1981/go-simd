package csum

//#cgo CFLAGS: -mavx -std=c17
//#include "sum.h"
import (
        "C"
)

import (
	"unsafe"
)

func SumFloat64C(buf []float64) (ret float64) {
	C.sum_float64_avx_intrinsics_c((*C.double)(unsafe.Pointer(&buf[0])), (C.ulong)(len(buf)), (*C.double)(unsafe.Pointer(&ret)))
	return ret
}
