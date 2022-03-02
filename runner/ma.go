package runner

import (
	"fmt"
	"log"

	"github.com/zhanglongx/Molokai/tsWrapper"
)

type ma struct {
	TsCode tsWrapper.TsCode

	// Upper indicts it cross from below or from above
	Upper bool

	// MA * (1 + Magnify)
	Magnify float64
}

func (m *ma) Valid() bool {
	return true
}

func (m *ma) Check() (r Result, err error) {
	r.TsCode = m.TsCode

	close, err := tsWrapper.RecentClose(m.TsCode)
	if err != nil {
		return Result{}, err
	}

	maV, err := tsWrapper.MA(m.TsCode)
	if err != nil {
		return Result{}, err
	}

	log.Printf("%s: ma %.2f close %.2f", m.TsCode, maV, close)

	if close < maV*(1+m.Magnify) && !m.Upper {
		r.IsShouldRemind = true
		r.Message = fmt.Sprintf("%.2f drop below MA %.2f", close, maV)
		return
	}

	if close > maV*(1+m.Magnify) && m.Upper {
		r.IsShouldRemind = true
		r.Message = fmt.Sprintf("%.2f exceed above MA %.2f", close, maV)
		return
	}

	return
}
