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

type Percent struct {
	Max float64 `json:"Max"`
	Min float64 `json:"Min"`
}

func (p *Percent) Run(symbol common.Symbol, date date.Date, cost float64,
	params *string) (bool, error) {

	if p.Max == 0.0 && p.Min == 0.0 {
		return false, errBadParams
	}

	p.Min = 11.0

	return true, nil
}

type XPercent Percent

// UnmarshalJSON Raise an Error in Go when Unmarshalling Unknown Fields from JSON
// It is important to note that If X was composed of real struct, the unmarshalling function would be
// recursive and break.
// see https://maori.geek.nz/golang-raise-error-if-unknown-field-in-json-with-exceptions-2b0caddecd1
func (p *Percent) UnmarshalJSON(data []byte) error {
	var m struct {
		XPercent
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields() // Force errors

	if err := dec.Decode(&m); err != nil {
		return err
	}

	*p = Percent(m.XPercent)
	return nil
}
