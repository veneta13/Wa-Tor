package WatorSim

import (
	"math/rand"
	"time"
)

func initBoard() [][]*creature {
	var board = make([][]*creature, *width)
	for i := range board {
		board[i] = make([]*creature, *height)
	}

	// create a new pseudorandom number generator
	prng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < *nFish; i++ {
		var x, y int
		for {
			x = prng.Intn(*width - 1)
			y = prng.Intn(*height - 1)

			if board[x][y] == nil {
				break
			}
		}

		board[x][y] = &creature{
			age:     rand.Intn(*fBreed),
			species: FISH,
		}
	}

	for i := 0; i < *nSharks; i++ {
		var x, y int
		for {
			x = prng.Intn(*width - 1)
			y = prng.Intn(*height - 1)

			if board[x][y] == nil {
				break
			}
		}

		board[x][y] = &creature{
			age:     rand.Intn(*sBreed),
			species: SHARK,
			health:  *starve,
		}
	}

	return board
}
