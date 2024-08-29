package capsis_ta_lib_csv

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"strconv"
)

func init() {
	zerolog.TimeFieldFormat = "15:04:05"
}

// Indicator line :t,h,l,c,BB_DOWN,BB_MID,BB_UP,spanA,spanB,tenkan,kijun,chiko
type Indicator struct {
	OpenTimeStr string  `csv:"t"`
	High        float64 `csv:"h"`
	Low         float64 `csv:"l"`
	Close       float64 `csv:"c"`
	BBUp        float64 `csv:"BB_DOWN"`
	BBMid       float64 `csv:"BB_MID"`
	BBDown      float64 `csv:"BB_UP"`
	ITenkan     float64 `csv:"tenkan"`
	IKijun      float64 `csv:"kijun"`
	ISpanA      float64 `csv:"spanA"`
	ISpanB      float64 `csv:"spanB"`
	IChiko      float64 `csv:"chiko"`
}

func NewIndicator(row []string) *Indicator {

	var err error
	var convHigh, convLow, convClose float64
	var convBBUp, convBBMid, convBBDown float64
	var convITenkan, convIKijun, convISpanA, convISpanB, convIChiko float64

	// HLC

	if len(row[1]) > 0 {
		convHigh, err = strconv.ParseFloat(row[1], 64)
		if err != nil {
			panic(err)
		}
	}

	if len(row[2]) > 0 {
		convLow, err = strconv.ParseFloat(row[2], 64)
		if err != nil {
			panic(err)
		}
	}

	if len(row[3]) > 0 {
		convClose, err = strconv.ParseFloat(row[3], 64)
		if err != nil {
			panic(err)
		}
	}

	// BB
	if len(row[4]) > 0 {
		convBBUp, err = strconv.ParseFloat(row[4], 64)
		if err != nil {
			panic(err)
		}
	}

	if len(row[5]) > 0 {
		convBBMid, err = strconv.ParseFloat(row[5], 64)
		if err != nil {
			panic(err)
		}
	}

	if len(row[6]) > 0 {
		convBBDown, err = strconv.ParseFloat(row[6], 64)
		if err != nil {
			panic(err)
		}
	}

	// Ichimoku

	if len(row[7]) > 0 {
		convISpanA, err = strconv.ParseFloat(row[7], 64)
		if err != nil {
			panic(err)
		}
	}

	if len(row[8]) > 0 {
		convISpanB, err = strconv.ParseFloat(row[8], 64)
		if err != nil {
			panic(err)
		}
	}

	if len(row[9]) > 0 {
		convITenkan, err = strconv.ParseFloat(row[9], 64)
		if err != nil {
			panic(err)
		}
	}

	if len(row[10]) > 0 {
		convIKijun, err = strconv.ParseFloat(row[10], 64)
		if err != nil {
			panic(err)
		}
	}

	if len(row[11]) > 0 {
		convIChiko, err = strconv.ParseFloat(row[11], 64)
		if err != nil {
			panic(err)
		}
	}

	return &Indicator{
		OpenTimeStr: row[0],
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
	}
}

func (i *Indicator) Log() {
	log.Info().
		Str("ot", i.OpenTimeStr).
		//Float64("h", i.High).
		//Float64("l", i.Low).
		//Float64("c", i.Close).
		//Float64("bbup", i.BBUp).
		//Float64("bbmid", i.BBMid).
		Float64("it", i.ITenkan).
		Float64("ik", i.IKijun).
		Float64("ia", i.ISpanA).
		Float64("ib", i.ISpanB).
		Float64("ic", i.IChiko).
		Msg("")
}
