package utils

import (
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"
)

// Function to copy images from source folder to destination folder
func CopyAllImagesFromDirectory(srcFolder, destFolder string) error {
	filesInDir, err := os.ReadDir(srcFolder)
	if err != nil {
		return err
	}

	if len(filesInDir) == 0 {
		return nil
	}

	return handleFilesInDirectory(filesInDir, srcFolder, destFolder)
}

func handleFilesInDirectory(filesInDir []fs.DirEntry, srcFolder, destFolder string) error {
	for _, file := range filesInDir {

		fileName := file.Name()

		if IsImageFile(fileName) {
			err := handleImageFile(srcFolder, destFolder, fileName)
			if err != nil {
				return err
			}
		}

		if file.IsDir() {
			err := handleDirFile(srcFolder, destFolder, fileName)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func handleDirFile(srcFolder, destFolder, fileName string) error {
	dirPath := path.Join(srcFolder, fileName)
	err := CopyAllImagesFromDirectory(dirPath, destFolder) // recursive call.
	return err
}

func handleImageFile(srcFolder, destFolder, fileName string) error {
	imgPath := path.Join(srcFolder, fileName)
	destImgPath := path.Join(destFolder, fileName)
	err := CopyFile(imgPath, destImgPath)
	return err
}

func IsImageFile(filename string) bool {
	imageExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp"}
	lowercaseFilename := strings.ToLower(filename)
	for _, ext := range imageExtensions {
		if strings.HasSuffix(lowercaseFilename, ext) {
			return true
		}
	}
	return false
}

func CopyFile(srcFile, dstFile string) error {
	log.Println("copying: " + srcFile + "to " + dstFile)
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer out.Close()

	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer in.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}
