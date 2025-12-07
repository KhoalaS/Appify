package core

import (
	"image"
	"image/color"
	"image/png"
	"math"

	"os"

	"golang.org/x/image/draw"
)

func ReadPngImage(pathToImage string) (image.Image, error) {
	imageFile, err := os.Open(pathToImage)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(imageFile)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func ScalePngImage(source image.Image, resolution int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, resolution, resolution))
	draw.NearestNeighbor.Scale(dst, dst.Rect, source, source.Bounds(), draw.Over, nil)

	return dst
}

func CircleCrop(src image.Image) image.Image {
	bounds := src.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()
	if w != h {
		panic("CircleCrop: image must be square")
	}

	dst := image.NewRGBA(bounds)
	r := float64(w) / 2
	cx := r
	cy := r

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			// Distance from center
			dx := float64(x) - cx + 0.5
			dy := float64(y) - cy + 0.5
			dist := math.Sqrt(dx*dx + dy*dy)

			if dist <= r { // inside circle
				dst.Set(x, y, src.At(x, y))
			} else { // outside circle = transparent
				dst.Set(x, y, color.RGBA{0, 0, 0, 0})
			}
		}
	}

	return dst
}

func WriteImage(source image.Image, filename string) error {
	outfile, err := os.Create(filename)
	if err != nil {
		return err
	}

	err = png.Encode(outfile, source)
	if err != nil {
		return err
	}
	return nil
}
