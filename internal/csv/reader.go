package capsis_ta_lib_csv

import (
	"encoding/csv"
	"os"
)

func ReadKlineCsv(csvPath string) []*Kline {
	var err error
	var file *os.File
	var res []*Kline
	var data [][]string

	// Open the CSV file
	file, err = os.Open(csvPath)
	if err != nil {
		panic(err)
	}

	// Read the CSV data
	reader := csv.NewReader(file)
	data, err = reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// load CSV data
	for i, row := range data {
		// skip header
		if i > 0 {
			newE := NewKline(row)
			res = append(res, newE)
		}
	}

	// close
	err = file.Close()
	if err != nil {
		panic(err)
	}

	return res
}

func ReadIndicatorCsv(csvPath string) []*Indicator {
	var err error
	var file *os.File
	var res []*Indicator
	var data [][]string

	// Open the CSV file
	file, err = os.Open(csvPath)
	if err != nil {
		panic(err)
	}

	// Read the CSV data
	reader := csv.NewReader(file)
	data, err = reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// load CSV data
	for i, row := range data {
		// skip header
		if i > 0 {
			newE := NewIndicator(row)
			res = append(res, newE)
		}
	}

	// close
	err = file.Close()
	if err != nil {
		panic(err)
	}

	return res
}
