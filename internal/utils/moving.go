package capsis_ta_utils

import (
	"errors"
	"fmt"
	"math"
)

func MovingMax(in []float64, period int) ([]float64, error) {

	if len(in) < period {
		return nil, errors.New(fmt.Sprintf("MovingMax: len(in): %d < period: %d", len(in), period))
	}

	var res = make([]float64, len(in))

	movingMax := math.Inf(-1)
	for i := 0; i < len(in); i++ {

		if i < period-1 {
			// init
			if in[i] > movingMax {
				movingMax = in[i]
			}
			res[i] = -1

		} else {
			// slide
			if in[i] > movingMax {
				movingMax = in[i]
			}

			res[i] = movingMax
		}
	}

	return res, nil
}

func MovingMin(in []float64, period int) ([]float64, error) {

	if len(in) < period {
		return nil, errors.New(fmt.Sprintf("MovingMin: len(in): %d < period: %d", len(in), period))
	}

	var res = make([]float64, len(in))

	movingMin := math.Inf(1)
	for i := 0; i < len(in); i++ {

		if i < period-1 {
			// init
			if in[i] < movingMin {
				movingMin = in[i]
			}
			res[i] = -1

		} else {
			// slide
			if in[i] < movingMin {
				movingMin = in[i]
			}

			res[i] = movingMin
		}
	}

	return res, nil
}
