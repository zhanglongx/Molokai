package runner

import (
	"fmt"
	"log"

	"github.com/zhanglongx/Molokai/tsWrapper"
)

type ma struct {
	Symbol string

	Upper bool
}

func (m *ma) Valid() bool {
	return true
}

func (m *ma) Check() (r Result, err error) {
	close, err := tsWrapper.RecentClose(m.Symbol)
	if err != nil {
		return Result{}, err
	}

	maV, err := tsWrapper.MA(m.Symbol)
	if err != nil {
		return Result{}, err
	}

	log.Printf("%s: ma %.2f close %.2f", m.Symbol, maV, close)

	if close < maV && !m.Upper {
		r.IsShouldRemind = true
		r.Message = fmt.Sprintf("%.2f drop below MA %.2f", close, maV)
		return
	}

	if close > maV && m.Upper {
		r.IsShouldRemind = true
		r.Message = fmt.Sprintf("%.2f exceed above MA %.2f", close, maV)
		return
	}

	return
}
