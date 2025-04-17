package demos

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name string
	Breed string
}


type Gorrila struct {
	Name string
	Color string
	NunberOfTheet int
}

func Do_myInterface(){
	dog := Dog {
		Name: "Samson",
		Breed: "German Shepherd",
	}

	PrintInfo(&dog)

	gorilla := Gorrila{
		Name: "Jock",
		Color: "Blak",
		NunberOfTheet: 38,
	}

	PrintInfo(&gorilla)
}

func PrintInfo (a Animal){
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "of legs")
}

func (d *Dog) Says() string{
	return "Woof"
}

func (d *Dog) NumberOfLegs() int{
	return 4
}

func (d *Gorrila) Says() string{
	return "Ugh"
}

func (d *Gorrila) NumberOfLegs() int{
	return 2
}