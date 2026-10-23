package main

import "fmt"

const (
	defaultMinTemperature = 15
	defaultMaxTemperature = 30
)

type Department struct {
	min int
	max int
}

func NewDepartment() Department {
	return Department{
		min: defaultMinTemperature,
		max: defaultMaxTemperature,
	}
}

func (d *Department) AddTemperature(sign string, temp int) int {
	if sign == ">=" {
		if temp > d.min {
			d.min = temp
		}
	} else if sign == "<=" {
		if temp < d.max {
			d.max = temp
		}
	}

	if d.min > d.max {
		return -1
	}

	return d.min
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	for i := 0; i < n; i++ {
		var staff int
		if _, err := fmt.Scan(&staff); err != nil {
			return
		}

		department := NewDepartment()

		for j := 0; j < staff; j++ {
			var sign string
			var temp int

			if _, err := fmt.Scan(&sign, &temp); err != nil {
				return
			}

			fmt.Println(department.AddTemperature(sign, temp))
		}
	}
}
