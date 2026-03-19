package entity

import "math"

type Temperature struct {
	C float64 `json:"temp_C"`
	F float64 `json:"temp_F"`
	K float64 `json:"temp_K"`
}

func round(num float64) float64 {
	return math.Round(num*100) / 100
}

func NewTemperature(c float64) Temperature {
	return Temperature{
		C: round(c),
		F: round(c*1.8 + 32),
		K: round(c + 273),
	}
}
