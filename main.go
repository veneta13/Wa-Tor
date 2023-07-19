package main

import (
	"Wa-Tor/WatorSim"
	"flag"
)

func init() {
	flag.IntVar(&WatorSim.InitFishCount, "fc", 10000, "Initial number of fish")
	flag.IntVar(&WatorSim.InitSharksCount, "sc", 400, "Initial number of sharks")
	flag.IntVar(&WatorSim.BreedFish, "fb", 50, "Length of fish breeding cycle")
	flag.IntVar(&WatorSim.BreedSharks, "sb", 70, "Length of shark breeding cycle")
	flag.IntVar(&WatorSim.Starve, "s", 150, "Chronons until starvation")
	flag.IntVar(&WatorSim.Width, "w", 1000, "Width of the board")
	flag.IntVar(&WatorSim.Height, "h", 1000, "Height of the board")

	flag.Parse()
}

func main() {

}
