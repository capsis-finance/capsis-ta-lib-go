package capsis_ta_utils_test

import (
	"fmt"
	capsis_ta_utils "github.com/capsian/capsis-ta-lib-go/internal/utils"
	"testing"
)

var (
	DATA1     = []float64{.01, .02, .03, .04, .05, .06, .07, .08, .09, .10, .11, .12, .13, .14, .15, .16, .17, .18, .19, .20, .21, .22, .23, .24}
	DATA1MAX9 = []float64{.0, .0, .0, .0, .0, .0, .0, .0, .09, .10, .11, .12, .13, .14, .15, .16, .17, .18, .19, .20, .21, .22, .23, .24}
	DATA1MIN9 = []float64{.01, .01, .01, .01, .01, .01, .01, .01, .01, .01, .01, .01, .01, .01, .01, .01, .01, .01, .01, .01, .01, .01, .01, .01}
	DATA2     = []float64{.01, .05, .08, .049, .0852, .056, .070, .080, .0999, .1520}
	DATA2MAX3 = []float64{0, 0, 0.08, 0.08, 0.0852, 0.0852, 0.0852, 0.0852, 0.0999, 0.152}
)

func TestMovingMax(t *testing.T) {

	res, err := capsis_ta_utils.MovingMax(DATA1, 9)
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < len(res); i++ {
		if res[i] != DATA1MAX9[i] {
			t.Errorf("i:%d, res[%d]:%f, DATA1MAX9[%d]:%f", i, i, res[i], i, DATA1MAX9[i])
		}
	}

	fmt.Println(res)

	res, err = capsis_ta_utils.MovingMax(DATA2, 3)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)

}

func TestMovingMin(t *testing.T) {

	res, err := capsis_ta_utils.MovingMin(DATA1, 9)
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < len(res); i++ {
		if res[i] != DATA1MIN9[i] {
			t.Errorf("i:%d, res[%d]:%f, MIN9[%d]:%f", i, i, res[i], i, DATA1MIN9[i])
		}
	}

	fmt.Println(res)
}
