package external_package

import (
	"fmt"
	"github.com/google/uuid"
)

func UUIDDemo() {
	for i := 0; i < 5; i++ {
		fmt.Println(uuid.New().String())
	}
}
