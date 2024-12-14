package error_package

import (
	"fmt"
	"time"
)

func Error() {
	err := fmt.Errorf("error occurred at: %v", time.Now())
	fmt.Println("An error happened", err)
}

func Boom() error {
	return fmt.Errorf("boom")
}
