// Copyright 2020 Longxiao Zhang <zhanglongx@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a GPLv3-style
// license that can be found in the LICENSE file.

package runner

import (
	"bytes"
	"encoding/json"

	"github.com/fxtlabs/date"
	"github.com/zhanglongx/Molokai/common"
)

type MinMax struct {
	Max float64 `json:"Max"`
	Min float64 `json:"Min"`
}

func (m *MinMax) Run(symbol common.Symbol, date date.Date, cost float64) (bool, error) {

	if m.Max == 0.0 && m.Min == 0.0 {
		return false, errBadParams
	}

	return true, nil
}

type XMinMax MinMax

// UnmarshalJSON Raise an Error in Go when Unmarshalling Unknown Fields from JSON
// It is important to note that If X was composed of real struct, the unmarshalling function would be
// recursive and break.
// see https://maori.geek.nz/golang-raise-error-if-unknown-field-in-json-with-exceptions-2b0caddecd1
func (m *MinMax) UnmarshalJSON(data []byte) error {
	var xm struct {
		XMinMax
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields() // Force errors

	if err := dec.Decode(&xm); err != nil {
		return err
	}

	*m = MinMax(xm.XMinMax)
	return nil
}
