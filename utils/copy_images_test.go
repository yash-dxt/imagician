package utils_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"yashandstuff.com/imagician/utils"
)

func TestIsImageFile(t *testing.T) {
	imageFiles := []string{
		"image1.jpg",
		"image2.png",
		"document.pdf", // Non-image file
		"image3.jpeg",
		"image4.bmp",
		"image5.gif",
		"image6.doc", // Non-image file
	}

	expectedImageExtensions := []bool{true, true, false, true, true, true, false}

	for i, filename := range imageFiles {
		isImage := utils.IsImageFile(filename)
		expectedIsImage := expectedImageExtensions[i]
		if isImage != expectedIsImage {
			t.Errorf("For file %s, expected isImage=%v, but got isImage=%v", filename, expectedIsImage, isImage)
		}
	}
}

func TestCopyFile(t *testing.T) {
	srcFile := "../test_files/dir1/cat.jpeg"
	destFile := "../test_files/dir2/cloned_cat.jpeg"

	// test that file doesn't exist.
	testNonExistenceOfFile(t, destFile)

	// copy file and test its existence
	err := utils.CopyFile(srcFile, destFile)
	assert.NoError(t, err)
	testExistenceOfFile(t, destFile)

	// delete file & check if deleted correctly.
	err = os.Remove(destFile)
	assert.NoError(t, err)
	testNonExistenceOfFile(t, destFile)
}

func TestCopyAllImagesFromDirectory(t *testing.T) {
	srcFile := "../test_files/cloned_from/"
	destFile := "../test_files/cloned_to/"

	err := utils.CopyAllImagesFromDirectory(srcFile, destFile)
	assert.NoError(t, err)

	testExistenceOfFile(t, destFile+"cat.jpeg")
	os.Remove(destFile + "cat.jpeg")
}

func testNonExistenceOfFile(t *testing.T, fileLocation string) {
	nonExistantFile, err := os.ReadFile(fileLocation)
	assert.Error(t, err)
	assert.Nil(t, nonExistantFile)
}

func testExistenceOfFile(t *testing.T, fileLocation string) {
	clonedFile, err := os.ReadFile(fileLocation)
	assert.NoError(t, err)
	assert.NotNil(t, clonedFile)
}
