package point

// #cgo CXXFLAGS: -std=c++17
// #cgo CFLAGS: -std=c11
// #cgo LDFLAGS: -static
// #include "cpoint.h"
// #include <stdlib.h>
import "C"
import (
	"log"
	"runtime"
	"unsafe"
)

//export Log
func Log(level int, msg *C.char) {
	log.Printf("%d - %s", level, C.GoString(msg))

	C.free(unsafe.Pointer(msg))
}

type Point struct {
	point *C.point_t
}

func NewPoint(pt *Point, x float64, y float64) *Point {
	if pt == nil {
		pt = &Point{}
		runtime.SetFinalizer(pt, DeletePoint)
	}

	pt.point = C.point_new(C.double(x), C.double(y))

	return pt
}

func (pt Point) X() float64 {
	return float64(C.point_x(pt.point))
}

func (pt Point) Y() float64 {
	return float64(C.point_y(pt.point))
}

func DeletePoint(pt *Point) {
	C.point_delete(pt.point)
}

func Distance(p *Point, q *Point) float64 {
	return float64(C.point_distance(p.point, q.point))
}
