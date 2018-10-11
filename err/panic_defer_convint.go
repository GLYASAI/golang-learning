package main

import (
	"fmt"
	"math"
)

func main() {
	l := int64(15000)
	if i, err := IntFromInt64(l); err != nil {
		fmt.Printf("The conversion of %d to an int32 resulted in an error: %s", l, err.Error())
	} else {
		fmt.Printf("%d converted to an int32 is %d", l, i)
	}
	fmt.Println()
	l = int64(math.MaxInt32 + 15000)
	if i, err := IntFromInt64(l); err != nil {
		fmt.Printf("The conversion of %d to an int32 resulted in an error: %s", l, err.Error())
	} else {
		fmt.Printf("%d converted to an int32 is %d", l, i)
	}
}

func ConvertInt64ToInt(input int64) int {
	if input >= math.MinInt32 && input <= math.MaxInt32 {
		return int(input)
	}
	panic(fmt.Sprintf("%d is out of the int32 range", input))
}

func IntFromInt64(input int64) (i int, err error) {
	defer func() {
		if err := recover(); err != nil {
			err = fmt.Errorf("%v", err)
			i = 123
		}
	}()

	return ConvertInt64ToInt(input), nil
}
