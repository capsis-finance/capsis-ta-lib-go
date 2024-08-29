package capsis_ta_lib_csv

import (
	"github.com/rs/zerolog/log"
	"strconv"
	"time"
)

// Indicator line :t,h,l,c,BB_DOWN,BB_MID,BB_UP,spanA,spanB,tenkan,kijun,chiko
type Indicator struct {
	OpenTime int64   `csv:"t"`
	High     float64 `csv:"h"`
	Low      float64 `csv:"l"`
	Close    float64 `csv:"c"`
	BBUp     float64 `csv:"BB_DOWN"`
	BBMid    float64 `csv:"BB_MID"`
	BBDown   float64 `csv:"BB_UP"`
	ITenkan  float64 `csv:"tenkan"`
	IKijun   float64 `csv:"kijun"`
	ISpanA   float64 `csv:"spanA"`
	ISpanB   float64 `csv:"spanB"`
	IChiko   float64 `csv:"chiko"`

	OpenTimeStr string
}

func NewIndicator(row []string) *Indicator {

	log.Info().Interface("row", row).Msg("NewIndicator")

	// Time

	convOpenTime, err := strconv.ParseInt(row[0], 10, 64)
	if err != nil {
		panic(err)
	}

	// HLC

	convHigh, err := strconv.ParseFloat(row[1], 64)
	if err != nil {
		panic(err)
	}

	convLow, err := strconv.ParseFloat(row[2], 64)
	if err != nil {
		panic(err)
	}

	convClose, err := strconv.ParseFloat(row[3], 64)
	if err != nil {
		panic(err)
	}

	// BB

	convBBUp, err := strconv.ParseFloat(row[4], 64)
	if err != nil {
		panic(err)
	}

	convBBMid, err := strconv.ParseFloat(row[5], 64)
	if err != nil {
		panic(err)
	}

	convBBDown, err := strconv.ParseFloat(row[6], 64)
	if err != nil {
		panic(err)
	}

	// Ichimoku

	convISpanA, err := strconv.ParseFloat(row[7], 64)
	if err != nil {
		panic(err)
	}

	convISpanB, err := strconv.ParseFloat(row[8], 64)
	if err != nil {
		panic(err)
	}

	convITenkan, err := strconv.ParseFloat(row[9], 64)
	if err != nil {
		panic(err)
	}

	convIKijun, err := strconv.ParseFloat(row[10], 64)
	if err != nil {
		panic(err)
	}

	convIChiko, err := strconv.ParseFloat(row[11], 64)
	if err != nil {
		panic(err)
	}

	return &Indicator{
		OpenTime:    convOpenTime,
		High:        convHigh,
		Low:         convLow,
		Close:       convClose,
		BBUp:        convBBUp,
		BBMid:       convBBMid,
		BBDown:      convBBDown,
		ITenkan:     convITenkan,
		IKijun:      convIKijun,
		ISpanA:      convISpanA,
		ISpanB:      convISpanB,
		IChiko:      convIChiko,
		OpenTimeStr: time.UnixMilli(convOpenTime).Format("2006-01-02 15:04:05"),
	}
}

func (i *Indicator) Log() {
	log.Info().
		Str("ot", i.OpenTimeStr).
		Float64("h", i.High).
		Float64("l", i.Low).
		Float64("c", i.Close).
		Float64("bbup", i.BBUp).
		Float64("bbmid", i.BBMid).
		Float64("it", i.ITenkan).
		Float64("ik", i.IKijun).
		Float64("ia", i.ISpanA).
		Float64("ib", i.ISpanB).
		Float64("ic", i.IChiko).
		Msg("")
}
