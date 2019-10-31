package main

import (
	"fmt"

	"github.com/bxcodec/slides/devfest/code/workshop/oop"
)

// START OMIT

func main() {
	pete := oop.Person{
		Name: "Peter"}
	askForName(pete)

	optimusPrime := oop.HumanLikeAlien{
		Eman: "Optimus Prime"}
	askForName(optimusPrime)
}

func askForName(h oop.Human) {
	fmt.Println("What's your name?")
	h.SayName()
}

// END OMIT
