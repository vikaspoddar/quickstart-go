package os

import (
	"fmt"
	"log"
	"os"
)

func OsDemo() {
	file, err := os.Open("os/example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = file.Close()
	if err != nil {
		log.Printf("issue while closing the file: %v", err)
	}
	fmt.Println(file.Name())
}
