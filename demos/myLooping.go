package demos

import "log"

func Do_looping() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}
}

func Do_looping_string() {

	animals := []string{"Dog", "Fish", "Cat", "Horse", "Cow"}

	for _, animal := range animals {
		log.Println(animal)
	}

}

func Do_looping_string2() {

	firtsLine := "Once upon a midnight dreary"

	for i, s := range firtsLine {
		log.Println(i, ":", s)
	}
}

func Do_looping_struct() {

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User

	users = append(users, User{"Kash", "Chet", "kash.chet@dev.com", 37})
	users = append(users, User{"Kman", "Ton", "kman.ton@dev.com", 38})
	users = append(users, User{"Bigus", "Dikus", "big.dick@dev.com", 40})
	users = append(users, User{"Bela", "Trisha", "bela.trix@dev.com", 25})

	for _, s := range users {
		log.Println(s.FirstName, s.LastName, s.Email, s.Age)
	}
}

func Do_looping_map() {

	myMap := make(map[string]string)

	myMap["dog"] = "Roxy"
	myMap["cat"] = "bocker"

	for animalType, animal := range myMap {
		log.Println(animalType, animal)
	}
}
