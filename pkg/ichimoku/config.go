package capsis_ta_ichimoku

type IchimokuConfig struct {
	TenkanPeriod int
	KijunPeriod  int
	SpanAPeriod  int
	SpanBPeriod  int
	ChikoPeriod  int

	SpanAProjectPeriod int
	SpanBProjectPeriod int

	initPeriod int
}
