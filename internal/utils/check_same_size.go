package capsis_ta_utils

func CheckSameSize(in1, in2 []float64) bool {
	return len(in1) == len(in2)
}

func CheckSameSize3(in1, in2, in3 []float64) bool {
	return CheckSameSize(in1, in2) && CheckSameSize(in1, in3)
}

func CheckSameSize4(in1, in2, in3, in4 []float64) bool {
	return CheckSameSize(in1, in2) && CheckSameSize(in1, in3) && CheckSameSize(in1, in4)
}
