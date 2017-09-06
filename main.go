package main

import (
	"fmt"
	"goms/goms"
	"math"
	"strconv"
)

func main() {
	val := "1000s"
	value, _ := goms.Ms(val)

	fmt.Printf("the value is :%f \n", value)
	fmt.Printf("the value is :%s \n", strconv.FormatInt(math.MaxInt64, 10))
	fmt.Printf("the value is :%d \n", len(strconv.FormatInt(math.MaxInt64, 10)))
}