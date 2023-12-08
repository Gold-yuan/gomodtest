package gomodtest

import (
	"fmt"
	"github.com/Gold-yuan/gomodtest/utils"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(utils.Sum(1, 2))
}

func SumDoubleB(a, b int) int {
	return a + b + b
}
