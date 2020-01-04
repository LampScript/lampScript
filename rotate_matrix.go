package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//[[1,2,3,4,5],[6,7,8,9,10],[11,12,13,14,15],[16,17,18,19,20],[21,22,23,24,25]]
func main() {
	//matrix := [][]int{
	//	[]int{1,2,3,4,5},
	//	[]int{6,7,8,9,10},
	//	[]int{11,12,13,14,15},
	//	[]int{16,17,18,19,20},
	//	[]int{21,22,23,24,25},
	//}
	//matrix := [][]int{
	//	[]int{1,2,3,4},
	//	[]int{5,6,7,8},
	//	[]int{9, 10,11,12},
	//	[]int{13,14,15,16},
	//}

	//rotate(matrix)
	//fmt.Println(matrix)
	StringToSlice("123")

}


// SliceToString slice to string with out data copy
func SliceToString(b []byte) (s string) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}

// StringToSlice string to slice with out data copy
func StringToSlice(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Println(reflect.TypeOf(pbytes).Kind().String())
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println(reflect.TypeOf(pstring).Kind().String())
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}


func rotate(matrix [][]int)  {

	if len(matrix) <= 1 {
		return
	}
	l := len(matrix)
	//转秩
	for i:=0;i< l;i++  {
		for j:=i;j<l ;j++{
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	fmt.Println(matrix)
	//reverse
	for i:=0;i<l;i++  {
		for j:=0;j<l/2 ;j++{
			matrix[i][j],matrix[i][l-j-1] = matrix[i][l-j-1], matrix[i][j]
		}
	}

}



