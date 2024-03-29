package WatorSim

import (
	"math/rand"
	"time"
)

func initBoard() [][]*creature {
	var board = make([][]*creature, Height)
	for i := range board {
		board[i] = make([]*creature, Width)
	}

	// create a new pseudorandom number generator
	prng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < InitFishCount; i++ {
		var x, y int
		for {
			x = prng.Intn(Width - 1)
			y = prng.Intn(Height - 1)

			if board[y][x] == nil {
				break
			}
		}

		board[y][x] = createFish()
	}

	for i := 0; i < InitSharksCount; i++ {
		var x, y int
		for {
			x = prng.Intn(Width - 1)
			y = prng.Intn(Height - 1)

			if board[y][x] == nil {
				break
			}
		}

		board[y][x] = createShark()
	}

	return board
}
