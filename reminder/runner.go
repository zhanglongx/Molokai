// Copyright 2020 Longxiao Zhang <zhanglongx@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a GPLv3-style
// license that can be found in the LICENSE file.

package runner

import (
	"encoding/json"
	"errors"

	"github.com/fxtlabs/date"
	"github.com/zhanglongx/Molokai/common"
	"github.com/zhanglongx/Molokai/tsWrapper"
)

type RunnerParams struct {
	Name   string      `json:"Name"`
	Params interface{} `json:"Params"`
}

type Runner interface {
	Run(symbol common.Symbol, date date.Date, cost float64) (bool, error)

	UnmarshalJSON(params []byte) error
}

var ts *tsWrapper.TsWrapper

var (
	errRunnerNotFound = errors.New("runner not found")
	errBadParams      = errors.New("bad params")
)

func RunnerInit() {
	if ts != nil {
		return
	}

	ts = &tsWrapper.TsWrapper{
		Token: "f4673f7862e73483c5e65cd9a036eedd39e72d484194a85dabcf958b",
	}

	ts.Init()
}

func RunnerRun(symbol common.Symbol, date date.Date, cost float64,
	params *string) (bool, error) {

	if ts == nil {
		RunnerInit()
	}

	var r RunnerParams
	if err := json.Unmarshal([]byte(*params), &r); err != nil {
		return false, err
	}

	var runner Runner
	switch r.Name {
	case "Percent":
		var Percent Percent
		runner = &Percent
	}

	if runner == nil {
		return false, errRunnerNotFound
	}

	// map[string]interface{} -> struct
	paramJson, err := json.Marshal(r.Params)
	if err != nil {
		return false, err
	}

	if err := runner.UnmarshalJSON(paramJson); err != nil {
		return false, err
	}

	changed, err := runner.Run(symbol, date, cost)
	if err != nil {
		return false, err
	}

	if changed {
		r.Params = runner
		paramJson, err = json.Marshal(r)
		if err != nil {
			return false, err
		}

		*params = string(paramJson)
	}

	return changed, nil
}
