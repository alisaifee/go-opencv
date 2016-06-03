package opencv

import (
	"testing"
	"github.com/lazywei/go-opencv/opencv"
	"io/ioutil"
)

func TestLoadImage(t *testing.T) {
	img := opencv.LoadImage("../images/pic5.png")
	if img.Width() != 400 || img.Height() != 300 {
		t.Fatal()
	}
}

func TestSaveImage(t *testing.T) {
	img := opencv.LoadImage("../images/pic5.png")
	tmpdir, _ := ioutil.TempDir("", "high-gui")
	filename := tmpdir + "/save_image.jpg"
	if 1 != opencv.SaveImage(filename, img, 0) {
		t.Fatal()
	}
}

func TestSaveImageInvalidDirectory(t *testing.T) {
	img := opencv.LoadImage("../images/pic5.png")
	tmpdir, _ := ioutil.TempDir("", "high-gui")
	filename := tmpdir + "/invalid/save_image.jpg"
	if 0 != opencv.SaveImage(filename, img, 0) {
		t.Fatal()
	}
}
