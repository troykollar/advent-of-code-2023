package read_input

import (
	"fmt"
	"os"
)

func ReadInput(name string) string {
	bytes, err := os.ReadFile(name)

	if err != nil {
		panic(fmt.Sprintf("An error occurred attempting to read %v: %v", name, err))
	}

	return string(bytes)
}
