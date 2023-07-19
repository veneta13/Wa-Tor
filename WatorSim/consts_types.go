package WatorSim

// constants

const (
	NORTH = iota
	SOUTH
	EAST
	WEST
)

const (
	FISH = iota
	SHARK
)

// flags

var (
	InitFishCount   int
	InitSharksCount int
	BreedFish       int
	BreedSharks     int
	Starve          int
	Width           int
	Height          int
)

// types

type Game struct{}

type coordinate struct {
	x, y int
}

type creature struct {
	age     int
	species int
	starve  int
}
