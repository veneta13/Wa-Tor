package WatorSim

import (
	"sync"
)

func columnPartition() []submatrix {
	partitionCount := ThreadCount

	var partitions []submatrix

	partSizeW := Width / partitionCount
	partRemW := Width % partitionCount

	fromX := 0
	for j := 0; j < partitionCount; j++ {
		toX := fromX + partSizeW - 1
		if j < partRemW {
			toX++
		}

		partitions = append(partitions, submatrix{
			fromX:      fromX,
			toX:        toX,
			fromY:      0,
			toY:        Height - 1,
			typeIsEven: false,
		})

		fromX = toX + 1
	}

	return partitions
}

func tickColumn(submatrixChan chan submatrix, board [][]*creature, waitGroup *sync.WaitGroup, mutex *sync.Mutex) {
	defer waitGroup.Done()

	for submatrix := range submatrixChan {
		mutex.Lock()
		for y := submatrix.fromY; y <= submatrix.toY; y++ {
			tickAnimal(board, submatrix.fromX, y)
		}
		mutex.Unlock()

		for x := submatrix.fromX + 1; x <= submatrix.toX; x++ {
			for y := submatrix.fromY; y <= submatrix.toY; y++ {
				tickAnimal(board, x, y)
			}
		}
	}
}

func runColumns(board [][]*creature, partitions []submatrix) {
	submatrixChan := make(chan submatrix, ThreadCount)
	waitGroup := sync.WaitGroup{}
	mutex := sync.Mutex{}

	for i := 0; i < ThreadCount; i++ {
		waitGroup.Add(1)
		go tickRow(submatrixChan, board, &waitGroup, &mutex)
	}

	for _, item := range partitions {
		submatrixChan <- item
	}

	close(submatrixChan)

	waitGroup.Wait()
}

func runColumn(board [][]*creature) {
	partitions := rowPartition()
	//var images []*image.Paletted

	for i := 0; i < MaxChronon; i++ {
		runColumns(board, partitions)
		//images = tickImage(images, board)
	}

	//createAnimation(images, "image.gif")
}

func CreateAndRunColumn() {
	board := initBoard()
	runColumn(board)
}
