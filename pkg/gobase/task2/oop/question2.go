package oop

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (employee Employee) PrintInfo() {
	fmt.Println("employeeID:", employee.EmployeeID)
	fmt.Println("name:", employee.Name)
	fmt.Println("age:", employee.Age)
}
