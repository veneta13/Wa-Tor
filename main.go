package main

import (
	"Wa-Tor/WatorSim"
	"flag"
)

func init() {
	flag.IntVar(&WatorSim.InitFishCount, "fc", 5000, "Initial number of fish")
	flag.IntVar(&WatorSim.InitSharksCount, "sc", 500, "Initial number of sharks")
	flag.IntVar(&WatorSim.BreedFish, "fb", 20, "Length of fish breeding cycle")
	flag.IntVar(&WatorSim.BreedSharks, "sb", 30, "Length of shark breeding cycle")
	flag.IntVar(&WatorSim.Starve, "s", 50, "Chronons until starvation")
	flag.IntVar(&WatorSim.Width, "w", 500, "Width of the board")
	flag.IntVar(&WatorSim.Height, "h", 500, "Height of the board")
	flag.IntVar(&WatorSim.ThreadCount, "t", 8, "Thread count")
	flag.IntVar(&WatorSim.MaxChronon, "mch", 1000, "Max chronons to run the simulation for")

	flag.Parse()
}

func main() {
	checkParams()
	WatorSim.CreateAndRunCheckboard()
}
