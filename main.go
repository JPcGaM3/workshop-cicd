package main

import (
	"fmt"

	"github.com/JPcGaM3/workshop-cicd/internal/calculator"
)
func main() {
	sum := calculator.Add(1, 2)
	fmt.Println("Hello world !!!", sum)
}
