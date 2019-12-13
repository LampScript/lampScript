package main

import (
	"fmt"
	"sync"
)

type Human struct {
	name string
}

type Father struct {
	h *Human
	age int
	sync.Mutex
}

//
//func main() {
//	i := new(int)
//	fmt.Printf("%p, %p \n", i, &i)
//	fmt.Println(reflect.TypeOf(i).Kind().String())
//	fmt.Println(reflect.TypeOf(&i).Kind().String())
//	s := new([]int)
//	sm := make([]int, 0)
//	fmt.Printf("%p, %p, %p \n", s, &s, *&s)
//	fmt.Println(reflect.TypeOf(s).Kind().String())
//	fmt.Println(reflect.TypeOf(&s).Kind().String())
//	fmt.Println(reflect.TypeOf(*&s).Kind().String())
//	fmt.Printf("%p, %v, %p \n", sm, &sm, *&sm)
//	fmt.Println(reflect.TypeOf(sm).Kind().String())
//	fmt.Println(reflect.TypeOf(&sm).Kind().String())
//	fmt.Println(reflect.TypeOf(*&sm).Kind().String())
//
//}

type canSort interface {
	GetInt() int
	Sort()
}

func main() {
	a := []int{43,23,56,7,4}
	qsort(a)
	fmt.Println(a)
}



func qsort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	qsort(data[:head])
	qsort(data[head+1:])
}

func addSlice(a map[int]string)  {
	a[1]="123"
}

func IsLocker(a sync.Locker)  {
	fmt.Println(a)
}