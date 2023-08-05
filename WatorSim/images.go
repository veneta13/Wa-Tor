package WatorSim

import (
	"image"
	"image/color"
	"image/gif"
	"os"
)

func createImage(board [][]*creature) *image.Paletted {
	img := image.NewPaletted(
		image.Rect(0, 0, Width, Height),
		color.Palette{waterColor, fishColor, sharkColor},
	)

	for i := 0; i < Width; i++ {
		for j := 0; j < Height; j++ {
			if board[i][j] == nil {
				img.Set(i, j, waterColor)
			} else {
				switch board[i][j].species {
				case FISH:
					img.Set(i, j, fishColor)
				case SHARK:
					img.Set(i, j, sharkColor)
				}
			}
		}
	}

	return img
}

func tickImage(images []*image.Paletted, board [][]*creature) []*image.Paletted {
	img := createImage(board)
	images = append(images, img)
	return images
}

func createAnimation(images []*image.Paletted, path string) {
	delay := make([]int, len(images))
	for i := range delay {
		delay[i] = 10
	}

	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delay,
	})
}
