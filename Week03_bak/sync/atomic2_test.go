package main
import (
"log"
"sync/atomic"
"testing"
"unsafe"
)

type X struct {
	v uint64
	x uint64
	a bool
	z uint64
	y uint32
}

/*
%p-输出指针值
unsafe.Offsetof-返回结构体中字段跟第一个地址的偏移量


 */
func TestAtomic(t *testing.T) {
	var x X
	log.Printf("x.a=%p, offset=%d, alig=%d", &x.a, unsafe.Offsetof(x.a), unsafe.Alignof(x.a))
	log.Printf("x.v=%p, offset=%d, alig=%d", &x.v, unsafe.Offsetof(x.v), unsafe.Alignof(x.v))
	log.Printf("x.x=%p, offset=%d, alig=%d", &x.x, unsafe.Offsetof(x.x), unsafe.Alignof(x.x))
	log.Printf("x.y=%p, offset=%d, alig=%d", &x.y, unsafe.Offsetof(x.y), unsafe.Alignof(x.y))
	log.Printf("x.z=%p, offset=%d, alig=%d", &x.z, unsafe.Offsetof(x.z), unsafe.Alignof(x.z))
	log.Printf("x.v=%p", &x.v)
	atomic.AddUint64(&x.z, 1) // panic
}

type Y struct {
	a bool
	X
}

func TestAtomicY(t *testing.T) {
	var y Y
	x := y.X
	atomic.AddUint64(&x.v, 1)
	atomic.AddUint64(&y.X.v, 1) // panic
}

type Y2 struct {
	X
	a bool
}

func TestAtomicY2(t *testing.T) {
	y := &Y2{}
	atomic.AddUint64(&y.X.v, 1)
}

type Temp struct {
	A byte
	B [2]byte
	C int64
}

func TestAtomicTemp(t *testing.T) {
	var x Temp
	log.Printf("sizof=%d", unsafe.Sizeof(x))
	log.Printf("x.A=%p, offset=%d, alig=%d", &x.A, unsafe.Offsetof(x.A), unsafe.Alignof(x.A))
	log.Printf("x.B=%p, offset=%d, alig=%d", &x.B, unsafe.Offsetof(x.B), unsafe.Alignof(x.B))
	log.Printf("x.C=%p, offset=%d, alig=%d", &x.C, unsafe.Offsetof(x.C), unsafe.Alignof(x.C))
}