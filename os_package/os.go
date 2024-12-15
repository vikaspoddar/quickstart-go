package os_package

import (
	"fmt"
	"os"
)

func OsDemo() {
	file, err := os.Open("os_package/example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	fmt.Println(file.Name())
}
