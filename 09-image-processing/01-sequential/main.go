package main

import (
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/disintegration/imaging"
)

func main() {
	start := time.Now()
	if err := walkFiles("img"); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Elapsed time: %v\n", time.Since(start))
}

func walkFiles(root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check that path is a regular file (no directory etc)
		if !info.Mode().IsRegular() {
			return nil
		}

		// Check that file is a jpeg, otherwise continue with next path
		contentType, _ := getFileContentType(path)
		if contentType != "image/jpeg" {
			fmt.Printf("File: %v is of type <%v> (should be <image/jpeg>), continue with next file\n", path, contentType)
			return nil
		}

		// Make the thumbnail, abort if error
		thumbnailImage, err := processImage(path)
		if err != nil {
			return err
		}

		// Save result, since last thing to do, always return err.
		// (Otherwise we would have an if-statement)
		err = saveThumbnail(path, thumbnailImage)
		return err
	})
}

func getFileContentType(file string) (string, error) {
	out, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer out.Close()

	buffer := make([]byte, 512)
	_, err = out.Read(buffer)
	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}

func processImage(path string) (*image.NRGBA, error) {
	srcImage, err := imaging.Open(path)
	if err != nil {
		return nil, err
	}

	thumbnailImage := imaging.Thumbnail(srcImage, 100, 100, imaging.Lanczos)
	return thumbnailImage, nil
}

func saveThumbnail(srcImagePath string, thumbnailImage *image.NRGBA) error {
	filename := filepath.Base(srcImagePath)
	dstImagePath := "thumbnail/" + filename

	err := imaging.Save(thumbnailImage, dstImagePath)
	if err != nil {
		return err
	}

	fmt.Printf("%s -> %s\n", srcImagePath, dstImagePath)
	return nil
}
