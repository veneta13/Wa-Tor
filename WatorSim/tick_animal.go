package WatorSim

import (
	"math"
	"math/rand"
)

func adjacent(x, y, direction int) coordinate {
	switch direction {
	case NORTH:
		return coordinate{
			x: x,
			y: int(math.Min(float64(y+1), float64(Height-1))),
		}
	case SOUTH:
		return coordinate{
			x: x,
			y: int(math.Max(float64(y-1), 0.0)),
		}
	case EAST:
		return coordinate{
			x: int(math.Min(float64(x+1), float64(Width-1))),
			y: y,
		}
	default: // WEST
		return coordinate{
			x: int(math.Max(float64(x-1), 0.0)),
			y: y,
		}
	}
}

func findEmptyAdjacent(board [][]*creature, x, y int) (int, int, int) {
	adjacentCoords := []coordinate{
		adjacent(x, y, NORTH),
		adjacent(x, y, SOUTH),
		adjacent(x, y, EAST),
		adjacent(x, y, WEST),
	}

	rand.Shuffle(
		len(adjacentCoords),
		func(i, j int) {
			adjacentCoords[i], adjacentCoords[j] = adjacentCoords[j], adjacentCoords[i]
		},
	)

	for _, tryCoord := range adjacentCoords {
		if tryCoord.x != x || tryCoord.y != y {
			if board[tryCoord.y][tryCoord.x] == nil {
				// return 1 - empty cell
				return 1, tryCoord.x, tryCoord.y
			}
			if board[y][x].species == SHARK && board[tryCoord.y][tryCoord.x].species == FISH {
				// return 2 - fish available
				return 2, tryCoord.x, tryCoord.y
			}
		}
	}
	return 0, 0, 0
}

func tickAnimal(board [][]*creature, x int, y int) {
	if board[y][x] == nil {
		return
	}

	board[y][x].age++
	board[y][x].starve--
	if board[y][x].starve <= 0 {
		board[y][x] = nil
		return
	}

	status, newX, newY := findEmptyAdjacent(board, x, y)

	if status == 0 {
		return
	}

	board[newY][newX] = board[y][x]

	if board[newY][newX].species == FISH {
		if board[newY][newX].age == BreedFish {
			board[y][x] = createFish()
		} else {
			board[y][x] = nil
		}
	} else { // the creature is a shark
		if board[newY][newX].age == BreedSharks {
			board[y][x] = createShark()
		} else {
			board[y][x] = nil
		}
	}

	if status == 2 {
		board[newY][newX].starve = Starve
	}
}
