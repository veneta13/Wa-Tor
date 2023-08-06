package WatorSim

import (
	"image"
	"sync"
)

func rowPartition() []submatrix {
	partitionCount := ThreadCount

	var partitions []submatrix

	partSizeH := Height / partitionCount
	partRemH := Height % partitionCount

	fromY := 0
	for j := 0; j < partitionCount; j++ {
		toY := fromY + partSizeH - 1
		if j < partRemH {
			toY++
		}

		partitions = append(partitions, submatrix{
			fromX:      0,
			toX:        Width - 1,
			fromY:      fromY,
			toY:        toY,
			typeIsEven: false,
		})

		fromY = toY + 1
	}

	return partitions
}

func tickRow(submatrixChan chan submatrix, board *[][]*creature, waitGroup *sync.WaitGroup, mutex *sync.Mutex) {
	defer waitGroup.Done()

	for submatrix := range submatrixChan {
		mutex.Lock()
		for x := submatrix.fromX; x <= submatrix.toX; x++ {
			tickAnimal(board, x, submatrix.fromY)
		}
		mutex.Unlock()

		for y := submatrix.fromY + 1; y <= submatrix.toY; y++ {
			for x := submatrix.fromX; x <= submatrix.toX; x++ {
				tickAnimal(board, x, y)
			}
		}
	}
}

func runRows(board *[][]*creature, partitions []submatrix) {
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

func runRow(board *[][]*creature) {
	partitions := rowPartition()
	var images []*image.Paletted

	for i := 0; i < MaxChronon; i++ {
		runRows(board, partitions)
		images = tickImage(images, board)
	}

	createAnimation(images, "image.gif")
}

func CreateAndRunRow() {
	board := initBoard()
	runRow(&board)
}
