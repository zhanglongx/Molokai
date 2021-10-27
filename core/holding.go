// Copyright 2020 Longxiao Zhang <zhanglongx@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a GPLv3-style
// license that can be found in the LICENSE file.

package core

import (
	"log"

	"github.com/zhanglongx/Molokai/runner"
	"github.com/zhanglongx/Molokai/tsWrapper"
)

type Holding struct {
	// Symbol. Same as TS ts_code
	Symbol string `yaml:"symbol"`

	// Runner and runner's param.
	Runners []runner.Runner `yaml:"runners"`
}

// Run launches all runners
func (h *Holding) Run(reminders Reminders) error {
	if h.Symbol == "" {
		log.Printf("symbol is nil")
		return nil
	}

	if len(h.Runners) == 0 {
		log.Printf("%s has no runner, skip", h.Symbol)
		return nil
	}

	for _, runner := range h.Runners {
		tsCode, err := tsWrapper.SymbolName(h.Symbol, "ts_code")
		if err != nil {
			log.Printf("symbol error: %v", err)
			continue
		}

		runner.TsCode = tsWrapper.TsCode(tsCode)

		m, err := runner.Run()
		if err != nil {
			log.Printf("run %s for runner %s failed: %v",
				runner.TsCode, runner.Name, err)
			continue
		}

		// TODO: use reminder list
		if m.IsShouldRemind {
			if reminders.IsTurnOn() {
				reminders.Fill(m)
			} else {
				log.Printf(m.Message)
			}
		}
	}

	// TODO: or launch reminders on all holdings?
	if err := reminders.Send(); err != nil {
		log.Printf("run reminder failed: %v", err)
	}

	return nil
}
