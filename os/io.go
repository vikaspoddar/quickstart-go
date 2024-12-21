package os

import (
	"io"
	"net/http"
	"os"
)

func IoDemo() {
	outputFile, _ := os.Create("os/novel.txt")
	defer outputFile.Close()
	response, _ := http.Get("https://www.gutenberg.org/cache/epub/74903/pg74903.txt")
	defer response.Body.Close()
	_, _ = io.Copy(outputFile, response.Body)

}
