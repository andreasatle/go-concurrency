package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	req, err := http.NewRequest("GET", "https://andcloud.io", nil)
	if err != nil {
		log.Fatal(err)
	}
	runWithTimeout(req, 100*time.Millisecond)
	runWithTimeout(req, 1000*time.Millisecond)
}

func runWithTimeout(req *http.Request, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(req.Context(), timeout)
	defer cancel()

	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
