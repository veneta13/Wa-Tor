package WatorSim

import (
	"math"
	"math/rand"
)

func createFish() *creature {
	return &creature{
		age:     rand.Intn(*breedFish),
		species: FISH,
		starve:  int(math.Inf(1)),
	}
}

func createShark() *creature {
	return &creature{
		age:     rand.Intn(*breedShark),
		species: SHARK,
		starve:  *starve,
	}
}
