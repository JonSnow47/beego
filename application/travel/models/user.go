package models

import (
	"fmt"
)

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func ID(num int) int {
	return num
}
