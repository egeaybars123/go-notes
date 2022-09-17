package main

import (
	"fmt"
)

type School interface {
	Welcome()
	StateDepartment()
}

type Student struct {
	name       string
	department string
	age        int
}

func (s Student) Welcome() {
	fmt.Printf("Welcome, %s to Bilkent University as a student! You are %d years old \n", s.name, s.age)
}

func (s Student) StateDepartment() {
	fmt.Printf("Your department is %s\n", s.department)
}

type Teacher struct {
	Student
	salary      int
	professions []string
}

type People struct {
	people []School
}

func (t Teacher) Welcome() {
	fmt.Printf("You are %s, and %d years old! Your salary is %d.\n", t.name, t.age, t.salary)
	fmt.Println("Your professions are:")
	for _, value := range t.professions {
		fmt.Println(value)
	}
}

func (p People) GiveInfo() {
	for _, v := range p.people {
		v.Welcome()
		v.StateDepartment()
	}
}

func main() {
	student1 := Student{"ege", "EEE", 20}
	teacher1 := Teacher{Student{"ahmet", "IE", 45}, 2000,
		[]string{"Blockchain", "Distributed Networks", "DeFi"}}

	bilkent := People{[]School{student1, teacher1}}
	bilkent.GiveInfo()
}
