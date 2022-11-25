package main

import "fmt"

type Employee struct {
    ID        int
    FirstName string
    LastName  string
    Address   string
}

func main() {
    employee := Employee{LastName: "Doe", FirstName: "John"}
    fmt.Println(employee)
    employeeCopy := &employee
    employeeCopy.FirstName = "David"
    fmt.Println(employee)
}


/*

type Employee struct {
    ID        int
    FirstName string
    LastName  string
    Address   string
}

*/


/*
package main

import "fmt"

type Person struct {
    ID        int
    FirstName string
    LastName  string
    Address   string
}

type Employee struct {
    Person
    ManagerID int
}

type Contractor struct {
    Person
    CompanyID int
}

func main() {
    employee := Employee{
        Person: Person{
            FirstName: "John",
        },
    }
    employee.LastName = "Doe"
    fmt.Println(employee.FirstName)
}

*/



/*
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    ID        int
    FirstName string `json:"name"`
    LastName  string
    Address   string `json:"address,omitempty"`
}

type Employee struct {
    Person
    ManagerID int
}

type Contractor struct {
    Person
    CompanyID int
}

func main() {
    employees := []Employee{
        Employee{
            Person: Person{
                LastName: "Doe", FirstName: "John",
            },
        },
        Employee{
            Person: Person{
                LastName: "Campbell", FirstName: "David",
            },
        },
    }

    data, _ := json.Marshal(employees)
    fmt.Printf("%s\n", data)

    var decoded []Employee
    json.Unmarshal(data, &decoded)
    fmt.Printf("%v", decoded)
}

*/