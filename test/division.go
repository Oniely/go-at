package main

import (
	"errors"
	"fmt"
)

func division() {
	var result, remainder, err = intDivision(2,2)

	switch {
	case err != nil:
		fmt.Println(err.Error())
	default:
		fmt.Println(result, remainder)
	}

}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error

	if denominator == 0 {
		err = errors.New("cannot devide by zero")
		return 0,0,err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, err
}
