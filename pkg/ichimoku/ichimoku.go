package capsis_ta_ichimoku

import (
	"errors"
	"fmt"
	capsis_indicator_utils "github.com/capsian/capsis-ta-lib-go/internal/utils"
	"time"
)

type PastValues struct {
	pastHigh  []float64
	pastLow   []float64
	pastClose []float64
}

type Ichimoku struct {
	Tenkan  []float64
	Kijun   []float64
	SpanA   []float64
	SpanB   []float64
	Chiko   []float64
	TsIndex []int64

	pastValues PastValues
	cfg        IchimokuConfig
	lastTsMs   int64
	intervalMs int64
}

func NewIchimoku(tenkanPeriod, kijunPeriod, spanAPeriod, spanBPeriod, chikoPeriod, spanAProjectPeriod, spanBProjectPeriod int, intervalMs int64) *Ichimoku {
	var res Ichimoku

	res.cfg = IchimokuConfig{
		TenkanPeriod:       tenkanPeriod,
		KijunPeriod:        kijunPeriod,
		SpanAPeriod:        spanAPeriod,
		SpanBPeriod:        spanBPeriod,
		ChikoPeriod:        chikoPeriod,
		SpanAProjectPeriod: spanAProjectPeriod,
		SpanBProjectPeriod: spanBProjectPeriod,
		initPeriod:         spanBPeriod + spanBProjectPeriod,
	}

	// init past values
	res.pastValues.pastHigh = make([]float64, 0)
	res.pastValues.pastLow = make([]float64, 0)
	res.pastValues.pastClose = make([]float64, 0)

	// init time index
	res.lastTsMs = 0
	res.intervalMs = intervalMs

	return &res
}

func (i *Ichimoku) Update(newHigh, newLow, newClose float64, newLastTsMs int64) (bool, error) {
	// init
	if i.lastTsMs == 0 {
		i.lastTsMs = newLastTsMs
	} else {
		// sanity check
		if i.lastTsMs+i.intervalMs != newLastTsMs {

			lastTsStr := time.UnixMilli(i.lastTsMs).Format("15:04:05")
			newLastTsStr := time.UnixMilli(newLastTsMs).Format("15:04:05")

			return false, errors.New(fmt.Sprintf("skipped klines, lastTsStr:%s, newLastTsStr:%s", lastTsStr, newLastTsStr))
		}

		// update
		i.lastTsMs = newLastTsMs
	}

	// init: accumulate
	if len(i.pastValues.pastHigh) != i.cfg.initPeriod {
		i.pastValues.pastHigh = append(i.pastValues.pastHigh, newHigh)
		i.pastValues.pastLow = append(i.pastValues.pastLow, newLow)
		i.pastValues.pastClose = append(i.pastValues.pastClose, newClose)
		i.TsIndex = append(i.TsIndex, newLastTsMs)
		return false, nil
	}

	// update: truncate
	i.pastValues.pastHigh = append(i.pastValues.pastHigh[1:], newHigh)
	i.pastValues.pastLow = append(i.pastValues.pastLow[1:], newLow)
	i.pastValues.pastClose = append(i.pastValues.pastClose[1:], newClose)
	i.TsIndex = append(i.TsIndex[1:], newLastTsMs)

	// process
	err := i.compute()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *Ichimoku) compute() error {

	var err error

	// tenkan
	i.Tenkan, err = i.computeTenkan()
	if err != nil {
		return err
	}

	// kijun
	i.Kijun, err = i.computeKijun()
	if err != nil {
		return err
	}

	// SpanA
	i.SpanA, err = i.computeSpanA()
	if err != nil {
		return err
	}

	// SpanB
	i.SpanB, err = i.computeSpanB()
	if err != nil {
		return err
	}

	// Chiko
	i.Chiko = i.computeChiko()

	return nil
}

func (i *Ichimoku) computeTenkan() ([]float64, error) {
	var err error
	var tenkanPeriodHigh, tenkanPeriodLow, tenkanPeriodSum, tenkanPeriodDiv []float64

	// Tenkan
	tenkanPeriodHigh, err = capsis_indicator_utils.MovingMax(i.pastValues.pastHigh, i.cfg.TenkanPeriod)
	if err != nil {
		return nil, err
	}
	tenkanPeriodLow, err = capsis_indicator_utils.MovingMin(i.pastValues.pastLow, i.cfg.TenkanPeriod)
	if err != nil {
		return nil, err
	}
	tenkanPeriodSum, err = capsis_indicator_utils.Add(tenkanPeriodHigh, tenkanPeriodLow)
	if err != nil {
		return nil, err
	}
	tenkanPeriodDiv = capsis_indicator_utils.DivideBy(tenkanPeriodSum, 2)

	return tenkanPeriodDiv, nil
}

func (i *Ichimoku) computeKijun() ([]float64, error) {
	var err error
	var kijunPeriodHigh, kijunPeriodLow, kijunPeriodSum, kijunPeriodDiv []float64

	kijunPeriodHigh, err = capsis_indicator_utils.MovingMax(i.pastValues.pastHigh, i.cfg.KijunPeriod)
	if err != nil {
		return nil, err
	}
	kijunPeriodLow, err = capsis_indicator_utils.MovingMin(i.pastValues.pastLow, i.cfg.KijunPeriod)
	if err != nil {
		return nil, err
	}
	kijunPeriodSum, err = capsis_indicator_utils.Add(kijunPeriodHigh, kijunPeriodLow)
	if err != nil {
		return nil, err
	}
	kijunPeriodDiv = capsis_indicator_utils.DivideBy(kijunPeriodSum, 2)

	return kijunPeriodDiv, nil
}

func (i *Ichimoku) computeSpanA() ([]float64, error) {
	var err error
	var spanASum, spanADiv []float64

	spanASum, err = capsis_indicator_utils.Add(i.Tenkan, i.Kijun)
	if err != nil {
		return nil, err
	}
	spanADiv = capsis_indicator_utils.DivideBy(spanASum, 2)

	// fill -1
	for j := 0; j < i.cfg.KijunPeriod; j++ {
		spanADiv[j] = -1
	}

	return spanADiv, nil
}

func (i *Ichimoku) computeSpanB() ([]float64, error) {
	var err error
	var spanBPeriodHigh, spanBPeriodLow, spanBPeriodSum, spanBPeriodDiv []float64

	spanBPeriodHigh, err = capsis_indicator_utils.MovingMax(i.pastValues.pastHigh, i.cfg.SpanBPeriod)
	if err != nil {
		return nil, err
	}
	spanBPeriodLow, err = capsis_indicator_utils.MovingMin(i.pastValues.pastLow, i.cfg.SpanBPeriod)
	if err != nil {
		return nil, err
	}
	spanBPeriodSum, err = capsis_indicator_utils.Add(spanBPeriodHigh, spanBPeriodLow)
	if err != nil {
		return nil, err
	}
	spanBPeriodDiv = capsis_indicator_utils.DivideBy(spanBPeriodSum, 2)

	return spanBPeriodDiv, nil
}

func (i *Ichimoku) computeChiko() []float64 {
	var chiko []float64

	chiko = capsis_indicator_utils.ShiftLeft(i.cfg.ChikoPeriod, i.pastValues.pastClose, -1)

	return chiko
}
