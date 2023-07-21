package main

import (
	"Wa-Tor/WatorSim"
	"log"
)

func checkParams() {
	creatureCount := WatorSim.InitFishCount + WatorSim.InitSharksCount
	boardSize := WatorSim.Width * WatorSim.Height

	if creatureCount > boardSize {
		log.Fatal("Board is too small!")
	}
}
