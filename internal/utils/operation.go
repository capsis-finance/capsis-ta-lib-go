package capsis_ta_utils

import (
	"errors"
	"fmt"
)

func Add(in1, in2 []float64) ([]float64, error) {
	if !CheckSameSize(in1, in2) {
		return in1, errors.New(fmt.Sprintf("in1:%d or in2:%d must be same size", len(in1), len(in2)))
	}

	result := make([]float64, len(in1))
	for i := 0; i < len(result); i++ {
		result[i] = in1[i] + in2[i]
	}

	return result, nil
}

func DivideBy(in1 []float64, div float64) []float64 {

	result := make([]float64, len(in1))
	for i := 0; i < len(in1); i++ {
		result[i] = in1[i] / div
	}

	return result
}

func ShiftLeft(period int, in1 []float64, fill float64) []float64 {

	result := make([]float64, len(in1))

	for i := 0; i < len(in1); i++ {
		if i+period+1 > len(in1) {
			result[i] = fill
		} else {
			result[i] = in1[i+period]
		}
	}

	return result
}
