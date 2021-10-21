package runner

import (
	"fmt"
	"math"
)

type minMax struct {
	Symbol string

	Min float64
	Max float64
}

func (m *minMax) Valid() bool {
	if m.Max == 0 {
		m.Max = math.MaxFloat64
	}

	return m.Max >= m.Min
}

func (m *minMax) Check() (r Result, err error) {
	now := 1.0
	if now < m.Min {
		r.IsShouldRemind = true
		r.Message = fmt.Sprintf("%.2f drop below %.2f", now, m.Min)
		return
	}

	if now > m.Max {
		r.IsShouldRemind = true
		r.Message = fmt.Sprintf("%.2f exceed %.2f", now, m.Max)
		return
	}

	return
}
