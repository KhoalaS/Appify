package core

import "testing"

func TestReadPngImage(t *testing.T) {
	_, err := ReadPngImage("./test_data/cs-black-000.png")
	if err != nil {
		t.Log(err)
		t.Error(err)
	}
}

func TestScaleImage(t *testing.T) {
	pngImage, err := ReadPngImage("./test_data/cs-black-000.png")
	if err != nil {
		t.Log(err)
		t.Error(err)
	}

	scaledImage := ScalePngImage(pngImage, 512)
	if scaledImage.Bounds().Dx() != 512 {
		t.Fail()
	}

	if scaledImage.Bounds().Dy() != 512 {
		t.Fail()
	}
}
