package capsis_ta_ichimoku_test

import (
	capsis_ta_lib_csv "github.com/capsis-finance/capsis-ta-lib-go/internal/csv"
	capsis_ta_ichimoku "github.com/capsis-finance/capsis-ta-lib-go/pkg/ichimoku"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"testing"
	"time"
)

func init() {
	zerolog.TimeFieldFormat = "15:04:05"
}

func TestNewIchimoku(t *testing.T) {

	// load expected indic
	inputIndicCsv := "../../test_data/test/test_indicator.csv"
	expectedIndic := capsis_ta_lib_csv.ReadIndicatorCsv(inputIndicCsv)

	// load kline
	inputKlineCsv := "../../test_data/test/test_kline.csv"
	kline := capsis_ta_lib_csv.ReadKlineCsv(inputKlineCsv)

	// init
	tenkanPeriod := 9
	kijunPeriod := 26
	spanBPeriod := 52
	chikoPeriod := 26
	spanAProjectPeriod := 26
	spanBProjectPeriod := 26
	intervalMs := time.Minute.Milliseconds()
	ichimoku := capsis_ta_ichimoku.NewIchimoku(tenkanPeriod, kijunPeriod, spanBPeriod, chikoPeriod, spanAProjectPeriod, spanBProjectPeriod, intervalMs)

	for i, _ := range expectedIndic {

		if i < len(kline) {
			k := kline[i]

			init, err := ichimoku.Update(k.High, k.Low, k.Close, k.OpenTime)
			if err != nil {
				log.Fatal().Err(err).Msg("")
			}

			if init {
				ichimoku.LogResult()
				// indic.Log()
			}
		}

	}
}
