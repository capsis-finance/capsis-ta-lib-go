package capsis_ta_lib_csv_test

import (
	capsis_ta_lib_csv "github.com/capsis-finance/capsis-ta-lib-go/internal/csv"
	"testing"
)

func TestReadKlineCsv(t *testing.T) {

	testFile := "../../" + "./test_data/test/test_kline.csv"

	res := capsis_ta_lib_csv.ReadKlineCsv(testFile)

	for _, rec := range res {
		rec.Log()
	}
}

func TestReadIndicatorCsv(t *testing.T) {

	testFile := "../../" + "./test_data/test/test_indicator.csv"

	res := capsis_ta_lib_csv.ReadIndicatorCsv(testFile)

	for _, rec := range res {
		rec.Log()
	}
}
