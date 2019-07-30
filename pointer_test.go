package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type foo struct {
	ID            string  `json:"_id"`
	Index         int     `json:"index"`
	GUID          string  `json:"guid"`
	IsActive      bool    `json:"isActive"`
	Balance       string  `json:"balance"`
	Picture       string  `json:"picture"`
	Age           int     `json:"age"`
	EyeColor      string  `json:"eyeColor"`
	Name          string  `json:"name"`
	Gender        string  `json:"gender"`
	Company       string  `json:"company"`
	Email         string  `json:"email"`
	Phone         string  `json:"phone"`
	Address       string  `json:"address"`
	About         string  `json:"about"`
	Registered    string  `json:"registered"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Greeting      string  `json:"greeting"`
	FavoriteFruit string  `json:"favoriteFruit"`
}

type bar struct {
	ID            string
	Index         int
	GUID          string
	IsActive      bool
	Balance       string
	Picture       string
	Age           int
	EyeColor      string
	Name          string
	Gender        string
	Company       string
	Email         string
	Phone         string
	Address       string
	About         string
	Registered    string
	Latitude      float64
	Longitude     float64
	Greeting      string
	FavoriteFruit string
}

var input foo

func init() {
	err := json.Unmarshal([]byte(`{
    "_id": "5d2f4fcf76c35513af00d47e",
    "index": 1,
    "guid": "ed687a14-590b-4d81-b0cb-ddaa857874ee",
    "isActive": true,
    "balance": "$3,837.19",
    "picture": "http://placehold.it/32x32",
    "age": 28,
    "eyeColor": "green",
    "name": "Rochelle Espinoza",
    "gender": "female",
    "company": "PARLEYNET",
    "email": "rochelleespinoza@parleynet.com",
    "phone": "+1 (969) 445-3766",
    "address": "956 Little Street, Jugtown, District Of Columbia, 6396",
    "about": "Excepteur exercitation labore ut cupidatat laboris mollit ad qui minim aliquip nostrud anim adipisicing est. Nisi sunt duis occaecat aliquip est irure Lorem irure nulla tempor sit sunt. Eiusmod laboris ex est velit minim ut cillum sunt laborum labore ad sunt.\r\n",
    "registered": "2016-03-20T12:07:25 -00:00",
    "latitude": 61.471517,
    "longitude": 54.01596,
    "greeting": "Hello, Rochelle Espinoza!You have 9 unread messages.",
    "favoriteFruit": "banana"
  }`), &input)
	if err != nil {
		panic(err)
	}
}

func byPointer(in *foo) *bar {
	return &bar{
		ID:            in.ID,
		Address:       in.Address,
		Email:         in.Email,
		Index:         in.Index,
		Name:          in.Name,
		About:         in.About,
		Age:           in.Age,
		Balance:       in.Balance,
		Company:       in.Company,
		EyeColor:      in.EyeColor,
		FavoriteFruit: in.FavoriteFruit,
		Gender:        in.Gender,
		Greeting:      in.Greeting,
		GUID:          in.GUID,
		IsActive:      in.IsActive,
		Latitude:      in.Latitude,
		Longitude:     in.Longitude,
		Phone:         in.Phone,
		Picture:       in.Picture,
		Registered:    in.Registered,
	}
}

func byValue(in foo) bar {
	return bar{
		ID:            in.ID,
		Address:       in.Address,
		Email:         in.Email,
		Index:         in.Index,
		Name:          in.Name,
		About:         in.About,
		Age:           in.Age,
		Balance:       in.Balance,
		Company:       in.Company,
		EyeColor:      in.EyeColor,
		FavoriteFruit: in.FavoriteFruit,
		Gender:        in.Gender,
		Greeting:      in.Greeting,
		GUID:          in.GUID,
		IsActive:      in.IsActive,
		Latitude:      in.Latitude,
		Longitude:     in.Longitude,
		Phone:         in.Phone,
		Picture:       in.Picture,
		Registered:    in.Registered,
	}
}

var pointerResult *bar

func BenchmarkByPointer(b *testing.B) {
	var r *bar
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = byPointer(&input)
	}
	pointerResult = r
}

var valueResult bar


func BenchmarkByValue(b *testing.B) {
	var r bar
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = byValue(input)
	}
	valueResult = r
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


func TestQuickSort(t *testing.T) {
	array := []int{
		55, 94,87,1,4,64,
	}

	heapFix(array, 0)
	fmt.Println(array)
}




func heapFix(a []int, i int)  {
	if i > len(a) - 1 {
		return
	}
	l := 2 * i + 1
	r := 2 * i + 2
	if l < len(a) && a[i] < a[l] {
		a[i], a[l] = a[l], a[i]
	}
	if r < len(a) && a[r] > a[i] {
		a[i], a[r] = a[r], a[i]
	}

	heapFix(a, l)
	heapFix(a, r)

}

func quickSort(array []int)  {
	if len(array) < 2 {
		return
	}

	mid := array[0]
	head, tail := 0, len(array) - 1
	for i := 1; i <= tail ; {
		if array[i] > mid {
			fmt.Println(array)
			array[i], array[tail] = array[tail], array[i]
			tail--
		} else {
			fmt.Println(array)
			array[i], array[head] = array[head], array[i]
			i++
			head++
		}
	}

	quickSort(array[:head])
	quickSort(array[head+1:])
}


