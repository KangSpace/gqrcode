package imaging

import (
	"github.com/disintegration/imaging"
	"os"
	"testing"
)

var home = os.Getenv("HOME")
var gqrcodePath = home + "/Desktop/gqrcode/"
var currentPath, _ = os.Getwd()

// TestGrayscale Grayscale Test
func TestGrayscale(t *testing.T) {
	testImage := gqrcodePath + "/image/grayscale_test.jpg"
	destImage := gqrcodePath + "/image/grayscale_result.jpg"
	if srcImage, err := imaging.Open(testImage); err == nil {
		grayscaleImage := Grayscale(srcImage)
		if err := imaging.Save(grayscaleImage, destImage); err == nil {
			t.Logf("image grayscale success, result: %s", destImage)
		} else {
			t.Fatalf("image save error: %s", testImage)
		}
	} else {
		t.Fatalf("image open error: %s", testImage)
	}
}

// TestBinarization Test grayscale image to binarization image
func TestBinarization(t *testing.T) {
	//testImage := gqrcodePath + "/image/grayscale_result.jpg"
	testImage := gqrcodePath + "/image/grayscale_test2.jpg"
	destImage := gqrcodePath + "/image/grayscale_binarization_result.jpg"
	if srcImage, err := imaging.Open(testImage); err == nil {
		binarizationImage := Binarization(srcImage)
		if err := imaging.Save(binarizationImage, destImage); err == nil {
			t.Logf("image binarization success, result: %s", destImage)
		} else {
			t.Fatalf("image save error: %s", testImage)
		}
	} else {
		t.Fatalf("image open error: %s", testImage)
	}
}

// TestImageBinarization Test raw image to binarization image
func TestImageBinarization(t *testing.T) {

}
