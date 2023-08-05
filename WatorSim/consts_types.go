package WatorSim

import "image/color"

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
	ThreadCount     int
	MaxChronon      int
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

type submatrix struct {
	fromX      int
	toX        int
	fromY      int
	toY        int
	typeIsEven bool
}

// params
var fishColor = color.RGBA{R: 221, G: 233, B: 9, A: 1}
var sharkColor = color.RGBA{R: 231, G: 1, B: 39, A: 1}
var waterColor = color.RGBA{R: 37, G: 87, B: 218, A: 1}
