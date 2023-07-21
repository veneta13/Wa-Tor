package WatorSim

import (
	"math"
)

func checkboardPartition() []submatrix {
	partitionCount := int(math.Sqrt(float64(2 * ThreadCount)))
	if partitionCount%2 == 1 {
		partitionCount++
	}

	partitions := make([]submatrix, 2*ThreadCount)

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

			partitions[i*partitionCount+j] = submatrix{
				fromX:      fromX,
				toX:        toX,
				fromY:      fromY,
				toY:        toY,
				typeIsEven: i%2 == j%2,
			}

			fromY = toY + 1
		}

		fromX = toX + 1
	}

	return partitions
}
