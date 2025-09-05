package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight = 1.8
	var userWeigt float64 = 100
	var IMT = userWeigt / math.Pow (userHeight, 2)
	fmt.Print(IMT)
}
