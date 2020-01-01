package main

import "fmt"

type Runner interface {
	Run()
	Say()
}

type Person struct {
	Name string
	_age int
}

func (p Person) Run() {
	fmt.Printf("%s is running\n", p.Name)
}

func (p *Person) Say() {
	fmt.Printf("hello, %s", p.Name)
}

func main() {
	fmt.Println("666")
	fmt.Println(1<<5)
}