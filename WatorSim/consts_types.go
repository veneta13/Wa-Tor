package WatorSim

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

type Game struct{}

type coordinate struct {
	x, y int
}

type creature struct {
	age     int
	species int
	starve  int
}
