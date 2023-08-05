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
			if board[tryCoord.x][tryCoord.y] == nil {
				// return 1 - empty cell
				return 1, tryCoord.x, tryCoord.y
			}
			if board[x][y].species == SHARK && board[tryCoord.x][tryCoord.y].species == FISH {
				// return 2 - fish available
				return 2, tryCoord.x, tryCoord.y
			}
		}
	}
	return 0, 0, 0
}

func tickAnimal(board [][]*creature, x int, y int) {
	if board[x][y] == nil {
		return
	}

	board[x][y].age++
	board[x][y].starve--
	if board[x][y].starve <= 0 {
		board[x][y] = nil
		return
	}

	status, newX, newY := findEmptyAdjacent(board, x, y)

	if status == 0 {
		return
	}

	board[newX][newY] = board[x][y]

	if board[newX][newY].species == FISH {
		if board[newX][newY].age == BreedFish {
			board[x][y] = createFish()
		} else {
			board[x][y] = nil
		}
	} else { // the creature is a shark
		if board[newX][newY].age == BreedSharks {
			board[x][y] = createShark()
		} else {
			board[x][y] = nil
		}
	}

	if status == 2 {
		board[newX][newY].starve = Starve
	}
}
