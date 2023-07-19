package WatorSim

import (
	"math"
	"math/rand"
)

func createFish() *creature {
	return &creature{
		age:     rand.Intn(BreedFish),
		species: FISH,
		starve:  int(math.Inf(1)),
	}
}

func createShark() *creature {
	return &creature{
		age:     rand.Intn(BreedSharks),
		species: SHARK,
		starve:  Starve,
	}
}
