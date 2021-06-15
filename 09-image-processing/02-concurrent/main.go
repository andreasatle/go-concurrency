package main

import (
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/disintegration/imaging"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	maxGoRoutines = runtime.GOMAXPROCS(runtime.NumCPU())
}

var maxGoRoutines int

func main() {
	start := time.Now()
	if err := pipeline("img"); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Elapsed time: %v\n", time.Since(start))
}

func pipeline(root string) error {
	paths := walkFiles(root)

	var wg sync.WaitGroup
	wg.Add(maxGoRoutines)
	for i := 0; i < maxGoRoutines; i++ {
		go func() {
			defer wg.Done()
			for path := range paths {
				processFile(path)
			}
		}()
	}

	wg.Wait()
	return nil
}

func walkFiles(root string) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Printf("Skipping path: %s, Error: %v\n", path, err)
				return nil
			}
			if !info.Mode().IsRegular() {
				return nil
			}

			ch <- path
			return nil
		})

		// Placeholder for error handling
		if err != nil {
			return
		}
	}()

	return ch
}

func processFile(path string) error {

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
	if err != nil {
		log.Println("Error saving thumbnail:", path)
	}

	return nil
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

	fmt.Printf("Created thumbnail: %s -> %s\n", srcImagePath, dstImagePath)
	return nil
}
