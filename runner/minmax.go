package runner

import (
	"fmt"
	"math"
)

type minMax struct {
	Symbol string

	Min  float64
	Max  float64
	Cost float64
}

func (m *minMax) Valid() bool {
	if m.Max == 0 {
		m.Max = math.MaxFloat64
	}

	return m.Max >= m.Min
}

func (m *minMax) Check() (r Result, err error) {
	if m.Cost < m.Min {
		r.IsShouldRemind = true
		r.Message = fmt.Sprintf("%f drop below %f", m.Cost, m.Min)
		return
	}

	if m.Cost > m.Max {
		r.IsShouldRemind = true
		r.Message = fmt.Sprintf("%f exceed %f", m.Cost, m.Max)
	}

	return
}
