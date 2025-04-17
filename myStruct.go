package main

import (
	"log"
)

type MyStruct struct {
	FirstName string
}

func (m *MyStruct) printFirstName() string {
	return m.FirstName;
}


func do_myStruct(){

	var myVar MyStruct
	 myVar.FirstName = "Kash" 

	 myVar2 := MyStruct{
		FirstName: "John",
	 }

	 log.Println("myVar is ", myVar.printFirstName())

	 log.Println("myVar2 is ", myVar2.printFirstName())
}