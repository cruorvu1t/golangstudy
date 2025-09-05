package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight := 1.8
	userWeigt := 100
	IMT := float64 (userWeigt) / math.Pow (userHeight, 2)
	fmt.Print(IMT)
}
