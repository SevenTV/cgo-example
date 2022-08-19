package point

// #cgo CXXFLAGS: -std=c++17
// #cgo CFLAGS: -std=c11
// #cgo LDFLAGS: -static
// #include "cpoint.h"
import "C"
import (
	"log"
	"runtime"
	"sync"
)

//export Log
func Log(level int, msg *C.char) {
	log.Printf("%d - %s", level, C.GoString(msg))
}

type Point struct {
	*cPoint
}

type cPoint struct {
	point *C.point_t
	once  sync.Once
}

func NewPoint(pt *Point, x float64, y float64) *Point {
	if pt == nil {
		pt = &Point{}
	}

	pt.cPoint = &cPoint{}
	runtime.SetFinalizer(pt.cPoint, func(pt *cPoint) {
		pt.delete()
	})

	pt.point = C.point_new(C.double(x), C.double(y))

	return pt
}

func (pt Point) X() float64 {
	return float64(C.point_x(pt.point))
}

func (pt Point) Y() float64 {
	return float64(C.point_y(pt.point))
}

func (pt *cPoint) delete() {
	pt.once.Do(func() {
		C.point_delete(pt.point)
	})
}

func DeletePoint(pt *Point) {
	pt.delete()
}

func Distance(p *Point, q *Point) float64 {
	return float64(C.point_distance(p.point, q.point))
}
