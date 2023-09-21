package main

import (
	"fmt"
	"basics/person"
)

func main() {

	var p person.Person = person.Person{
		Name: "John Doe",
		Age:  30,
		Address: "Jakarta",
		CreditScoring: 700.54,
	}

	message, _ := p.SayHello()

	fmt.Println(person.SumIntOrFloat[int](p.Age, 13))
	fmt.Println(person.SumIntOrFloat[float32](p.CreditScoring, 13.5))

	fmt.Println(message)

}