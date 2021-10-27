package runner

import (
	"fmt"
	"log"
	"math"

	"github.com/zhanglongx/Molokai/tsWrapper"
)

type minMax struct {
	TsCode tsWrapper.TsCode

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
	now, err := tsWrapper.RecentClose(m.TsCode)
	if err != nil {
		return Result{}, err
	}

	log.Printf("%s: close %.2f", m.TsCode, now)

	if now < m.Min {
		r.IsShouldRemind = true
		r.Message = fmt.Sprintf("%.2f drop below %.2f", now, m.Min)
		return
	}

	if now > m.Max {
		r.IsShouldRemind = true
		r.Message = fmt.Sprintf("%.2f exceed above %.2f", now, m.Max)
		return
	}

	return
}
