package tsWrapper

import (
	"errors"
	"strings"

	"github.com/ShawnRong/tushare-go"
	"github.com/fxtlabs/date"
	"github.com/zhanglongx/Molokai/common"
)

type TsWrapper struct {
	Token string
	api   *tushare.TuShare
}

type Value struct {
	Close float64
	Date  date.Date
}

var (
	errEmptyData = errors.New("empty data, wrong token? not in trading day?")
	errType      = errors.New("type error")
)

func (t *TsWrapper) Init() {
	t.api = tushare.New(t.Token)
}

// AdjFactor https://tushare.pro/document/2?doc_id=28
// Only if date is trading days can get the data, if it is a non-trading day
// error will be returned
func (t *TsWrapper) AdjFactor(tsCode common.Symbol, date date.Date) (float64, error) {
	params := make(map[string]string)
	params["ts_code"] = string(tsCode)
	params["start_date"] = toNumeric(date)
	params["end_date"] = toNumeric(date)

	fields := []string{"adj_factor"}

	resp, err := t.api.AdjFactor(params, fields)
	if err != nil {
		return 1.0, err
	}

	if len(resp.Data.Items) == 0 {
		return 1.0, errEmptyData
	}

	adj, ok := resp.Data.Items[0][0].(float64)
	if !ok {
		return 1.0, errType
	}

	return adj, nil
}

func (t *TsWrapper) RecentClose(tsCode common.Symbol) (float64, error) {
	params := make(map[string]string)
	params["ts_code"] = string(tsCode)

	// fields := []string{"close"}

	// TODO:

	return 1.0, errEmptyData
}

func toNumeric(d date.Date) string {
	from := d.String()
	return strings.Replace(from, "-", "", -1)
}
