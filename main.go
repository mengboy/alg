package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//fmt.Println(math.Round(6.4))
	//fmt.Println(round(6.4))
	//
	//fmt.Println(math.Round(6.5))
	//fmt.Println(round(6.5))
	//
	//fmt.Println(math.Round(6.6))
	//fmt.Println(round(6.6))
	//fmt.Println(math.Round(6.45))
	//fmt.Println(round(6.45))
	//

	str := "0"

	n, err := strconv.ParseFloat(str, 64)
	fmt.Println(n, err)

}

func round(x float64) float64 {
	return math.Floor(x + 0.5)
}
