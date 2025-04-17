package demos

import "log"

func Do_conditional() {

	var isTrue bool

	isTrue = false

	if isTrue {
		log.Println("is True:", isTrue)
	} else {
		log.Println("is True:", isTrue)
	}

	cat := "cat2"

	if cat == "cat" {
		log.Println("cat is cat")
	} else {
		log.Println("cat is not cat")
	}

	myNum := 1
	isTrue = false

	if myNum > 99 && isTrue {
		log.Println("myNum greater than 99 and true")
	} else if myNum > 99 {
		log.Println("myNum is just greater than 99")
	} else {
		log.Println("myNum is not greater than 99 or false")
	}
}

func Do_switch_condition() {
	myVar := "dragon"

	switch myVar {
	case "cat":
		log.Println("Cat")

	case "dog":
		log.Println("Dog")

	case "fish":
		log.Println("Fish")

	default:
		log.Println("Something Else")
	}

}
