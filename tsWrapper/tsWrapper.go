package tsWrapper

import (
	"errors"
	"strings"

	"github.com/ShawnRong/tushare-go"
	"github.com/fxtlabs/date"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

var Token string

var api *tushare.TuShare

var (
	errTokenNotSet = errors.New("tushare token is not set")
	errEmptyData   = errors.New("empty data, wrong token? not in trading day?")
	errValue       = errors.New("value error")
)

func apiInit() error {
	if api == nil {
		if Token == "" {
			return errTokenNotSet
		}

		api = tushare.New(Token)
	}

	return nil
}

// Symbol returns the ts_code, the input can be either of
// ts_code, symbol, name. It Will get the ts_code value by
// querying the StockBasic api
func Symbol(s string, key string) (string, error) {
	if err := apiInit(); err != nil {
		return "", err
	}

	fields := []string{"ts_code", "symbol", "name"}

	resp, err := api.StockBasic(map[string]string{"exchange": "", "list_status": "L"},
		fields)
	if err != nil {
		return "", err
	}

	if len(resp.Data.Items) == 0 {
		return "", errEmptyData
	}

	index := -1
	for i := range fields {
		if fields[i] == key {
			index = i
		}
	}

	if index == -1 {
		return "", errValue
	}

	for _, row := range resp.Data.Items {
		for _, col := range row {
			if col.(string) == s {
				return row[index].(string), nil
			}
		}
	}

	return "", errEmptyData
}

// AdjFactor https://tushare.pro/document/2?doc_id=28
// Only if date is trading days can get the data, if it is a non-trading day
// error will be returned
func AdjFactor(tsCode string, date date.Date) (float64, error) {
	if err := apiInit(); err != nil {
		return 1.0, err
	}

	fields := []string{"adj_factor"}

	resp, err := api.AdjFactor(map[string]string{
		"ts_code":    tsCode,
		"start_date": toNumeric(date),
		"end_date":   toNumeric(date)}, fields)
	if err != nil {
		return 1.0, err
	}

	if len(resp.Data.Items) == 0 {
		return 1.0, errEmptyData
	}

	return resp.Data.Items[0][0].(float64), nil
}

// RecentClose returns the recent one close value
func RecentClose(tsCode string) (float64, error) {
	s, err := Close(tsCode)
	if err != nil {
		return 0.0, err
	}

	return s.Elem(0).Float(), nil
}

// Close return a series of close by quering DailyBasic API
func Close(tsCode string) (series.Series, error) {
	if err := apiInit(); err != nil {
		return series.Series{}, err
	}

	params := make(map[string]string)
	params["ts_code"] = string(tsCode)

	fields := []string{"close"}

	resp, err := api.DailyBasic(params, fields)
	if err != nil {
		return series.Series{}, err
	}

	if len(resp.Data.Items) == 0 {
		return series.Series{}, errEmptyData
	}

	df, err := array2DtoDf(resp.Data.Items, fields)
	if err != nil {
		return series.Series{}, err
	}

	return df.Col("close"), nil
}

// MA return a the default 60 period of Moving Average
// close by quering DailyBasic API. It *NOT* support
// AdjFactor now
func MA(tsCode string) (float64, error) {
	s, err := Close(tsCode)
	if err != nil {
		return 0.0, err
	}

	// TODO: adjFactor
	return s.Slice(0, 60).Mean(), nil
}

func toNumeric(d date.Date) string {
	from := d.String()
	return strings.Replace(from, "-", "", -1)
}

func array2DtoDf(a [][]interface{}, fields []string) (*dataframe.DataFrame, error) {
	result := []map[string]interface{}{}

	for _, row := range a {
		tmp := make(map[string]interface{})
		for i, col := range fields {
			tmp[col] = row[i]
		}
		result = append(result, tmp)
	}

	df := dataframe.LoadMaps(result)
	return &df, df.Error()
}
