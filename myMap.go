package main

import (
	"log"
	"sort"
)

func do_myMap(){
	myMap := make(map[string]string)

	myMap["dog"] = "Roxy"

	myMap["cat"] = "bocker"

	log.Println("mymap", myMap["dog"])

	log.Println("mymap", myMap["cat"])

	myMap2 := make(map[string]int)

	myMap2["First"] = 1
	myMap2["Second"] = 2

	log.Println("myMap2", myMap2["First"])
	log.Println("myMap2", myMap2["Second"])

	user := MyStruct{
		FirstName: "Kman",
	}

    // ------------------------------------------------------- //

	myMap3 := make(map[string]MyStruct)

	myMap3["user1"] = user

	log.Println("myMap3", myMap3["user1"].FirstName)


	// ------------------------------------------------------- //

	var mySlice []string

	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Gabe")
	mySlice = append(mySlice, "Jack")
	mySlice = append(mySlice, "Peeter")
	mySlice = append(mySlice, "North")


	log.Println(mySlice)

	sort.Strings(mySlice)

	log.Println(mySlice)


	myNumbers := []int {1,2,3,4,5,6,7,8,9}

	log.Println(myNumbers)

	log.Println(myNumbers[0:2])

	log.Println(myNumbers[6:9])

}