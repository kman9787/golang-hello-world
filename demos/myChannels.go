package demos

import (
	"fmt"

	"github.com/kman9787/golang-hello-world/helpers"
)

func Do_channels() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <- intChan

	fmt.Println("intChan", num)
}

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNember(numPool)
	intChan <- randomNumber
}