package capsis_ta_lib_csv

import (
	"github.com/rs/zerolog/log"
	"strconv"
	"time"
)

type Kline struct {
	OpenTime    int64
	Open        float64
	High        float64
	Low         float64
	Close       float64
	VolumeBase  float64
	VolumeQuote float64
	NbTrade     int64

	OpenTimeStr string
}

func NewKline(row []string) *Kline {

	convTime, err := strconv.ParseInt(row[0], 10, 64)
	if err != nil {
		panic(err)
	}

	convOpen, err := strconv.ParseFloat(row[1], 64)
	if err != nil {
		panic(err)
	}

	convHigh, err := strconv.ParseFloat(row[2], 64)
	if err != nil {
		panic(err)
	}

	convLow, err := strconv.ParseFloat(row[3], 64)
	if err != nil {
		panic(err)
	}

	convClose, err := strconv.ParseFloat(row[4], 64)
	if err != nil {
		panic(err)
	}

	convVolBase, err := strconv.ParseFloat(row[5], 64)
	if err != nil {
		panic(err)
	}

	convVolQuote, err := strconv.ParseFloat(row[6], 64)
	if err != nil {
		panic(err)
	}

	convNbTrade, err := strconv.ParseInt(row[7], 10, 64)
	if err != nil {
		panic(err)
	}

	return &Kline{
		OpenTime:    convTime,
		Open:        convOpen,
		High:        convHigh,
		Low:         convLow,
		Close:       convClose,
		VolumeBase:  convVolBase,
		VolumeQuote: convVolQuote,
		NbTrade:     convNbTrade,
		OpenTimeStr: time.UnixMilli(convTime).Format("2006-01-02 15:04:05"),
	}
}

func (k *Kline) Log() {
	log.Info().
		Str("ot", k.OpenTimeStr).
		Float64("h", k.High).
		Float64("l", k.Low).
		Float64("c", k.Close).
		Msg("")
}
