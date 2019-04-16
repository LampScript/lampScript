package test



import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)
var mu sync.Mutex
//go:noinline
func foo() {}
//go:noinline
func deferLockTwo() {
	mu.Lock()
	defer mu.Unlock()
	defer foo()
}
//go:noinline
func deferLock() {
	mu.Lock()
	defer mu.Unlock()
	foo()
}
//go:noinline
func deferLockClosure() {
	mu.Lock()
	defer func() { mu.Unlock() }()
	foo()
}
//go:noinline
func noDeferLock() {
	mu.Lock()
	mu.Unlock()
	foo()
}
func BenchmarkDeferLockTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deferLockTwo()
	}
}
func BenchmarkDeferLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deferLock()
	}
}
func BenchmarkDeferLockClosure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deferLockClosure()
	}
}
func BenchmarkNoDeferLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		noDeferLock()
	}
}




func foo1(i int)  {
	i++
}


func temp()  {
}

type FooObj interface {
	foo2(i int)
	temp2()
}

type name struct {

}

func (n *name) foo2(i int)  {
	i++
}

func (n *name) temp2()  {
}

//BenchmarkFuncInSide-4   	2000000000	         0.32 ns/op
func BenchmarkFuncInSide(b *testing.B) {
	n := new(name)
	for i := 0; i < b.N; i++ {
		n.foo2(i)
		n.temp2()
	}
}

//BenchmarkFuncOut-4   	2000000000	         0.31 ns/op
func BenchmarkFuncOut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo1(i)
		temp()
	}
}

func TestFuc(t *testing.T)  {
	a := func() {
		fmt.Println("aaa")
	}
	ppp(a)

}

func ppp(in interface{})  {
	fmt.Println(reflect.TypeOf(in).Kind())
}

