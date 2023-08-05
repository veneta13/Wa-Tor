package WatorSim

import (
	"image"
	"math"
	"sync"
)

var waitGroup sync.WaitGroup

func checkboardPartition() ([]submatrix, []submatrix) {
	partitionCount := int(math.Sqrt(float64(2 * ThreadCount)))
	if partitionCount%2 == 1 {
		partitionCount++
	}

	evenPartitions := make([]submatrix, ThreadCount)
	oddPartitions := make([]submatrix, ThreadCount)

	partSizeW := Width / partitionCount
	partRemW := Width % partitionCount
	partSizeH := Height / partitionCount
	partRemH := Height % partitionCount

	fromX := 0
	for i := 0; i < partitionCount; i++ {
		toX := fromX + partSizeW - 1
		if i < partRemW {
			toX++
		}

		fromY := 0
		for j := 0; j < partitionCount; j++ {
			toY := fromY + partSizeH - 1
			if j < partRemH {
				toY++
			}

			if i%2 == j%2 {
				evenPartitions[(i*partitionCount+j)/2] = submatrix{
					fromX:      fromX,
					toX:        toX,
					fromY:      fromY,
					toY:        toY,
					typeIsEven: true,
				}
			} else {
				oddPartitions[(i*partitionCount+j)/2] = submatrix{
					fromX:      fromX,
					toX:        toX,
					fromY:      fromY,
					toY:        toY,
					typeIsEven: false,
				}
			}

			fromY = toY + 1
		}

		fromX = toX + 1
	}

	return evenPartitions, oddPartitions
}

func tickCheckboard(submatrixChan chan submatrix, board [][]*creature) {
	defer waitGroup.Done()

	submatrix := <-submatrixChan
	for x := submatrix.fromX; x <= submatrix.toX; x++ {
		for y := submatrix.fromY; y <= submatrix.toY; y++ {
			tickAnimal(board, x, y)
		}
	}
}

func runHalf(board [][]*creature, partitions []submatrix) {
	submatrixChan := make(chan submatrix, ThreadCount)

	for i := 0; i < ThreadCount; i++ {
		waitGroup.Add(1)
		go tickCheckboard(submatrixChan, board)
	}

	for _, item := range partitions {
		submatrixChan <- item
	}

	close(submatrixChan)

	waitGroup.Wait()
}

func runCheckboard(board [][]*creature) {
	evenPartitions, oddPartitions := checkboardPartition()
	var images []*image.Paletted

	for i := 0; i < MaxChronon; i++ {
		runHalf(board, evenPartitions)
		runHalf(board, oddPartitions)
		images = tickImage(images, board)
	}

	createAnimation(images, "image.gif")
}

func CreateAndRunCheckboard() {
	board := initBoard()
	runCheckboard(board)
}
