package main

import (
	"Wa-Tor/WatorSim"
	"flag"
	"fmt"
	"runtime"
	"time"
)

func init() {
	flag.IntVar(&WatorSim.InitFishCount, "fc", 100000, "Initial number of fish")
	flag.IntVar(&WatorSim.InitSharksCount, "sc", 1000, "Initial number of sharks")
	flag.IntVar(&WatorSim.BreedFish, "fb", 20, "Length of fish breeding cycle")
	flag.IntVar(&WatorSim.BreedSharks, "sb", 30, "Length of shark breeding cycle")
	flag.IntVar(&WatorSim.Starve, "s", 50, "Chronons until starvation")
	flag.IntVar(&WatorSim.Width, "w", 2000, "Width of the board")
	flag.IntVar(&WatorSim.Height, "h", 2000, "Height of the board")
	flag.IntVar(&WatorSim.ThreadCount, "t", 4, "Thread count")
	flag.IntVar(&WatorSim.MaxChronon, "mch", 100, "Max chronons to run the simulation for")
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	checkParams()

	startTime := time.Now()
	WatorSim.CreateAndRunColumn()

	elapsed := time.Since(startTime)
	fmt.Printf("Total execution time: %s\n", elapsed)
}
