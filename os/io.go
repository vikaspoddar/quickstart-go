package os

import (
	"io"
	"log"
	"net/http"
	"os"
)

func IoDemo() {
	outputFile, err := os.Create("os/novel.txt")
	if err != nil {
		log.Printf("filed to create a file: %v", err)
		return
	}
	err = outputFile.Close()
	if err != nil {
		log.Printf("file not closed: %v", err)
	}
	response, err := http.Get("https://www.gutenberg.org/cache/epub/74903/pg74903.txt")
	if err != nil {
		log.Printf("failed to fetch novel: %v", err)
	}
	_, _ = io.Copy(outputFile, response.Body)

}
