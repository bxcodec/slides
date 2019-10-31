package main

import "github.com/bxcodec/slides/devfest/code/workshop/composition"

func main() {
	elonMusk := composition.AspiredEntrepreneur{
		composition.BusinessGuy{"life on Mars"},
		composition.TechieGuy{"reusable rockets"},
	}
	composition.Pitch(elonMusk)
}
