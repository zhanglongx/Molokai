// Copyright 2020 Longxiao Zhang <zhanglongx@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a GPLv3-style
// license that can be found in the LICENSE file.

package core

import (
	"fmt"
	"log"

	"github.com/zhanglongx/Molokai/runner"
)

type Holding struct {
	// Symbol. Same as TS ts_code
	Symbol string `yaml:"symbol"`

	// Runner and runner's param.
	Runners []runner.Runner `yaml:"runners"`
}

// Runholdings launches runners on each holding
func RunHoldings(holdings []Holding) error {
	if len(holdings) == 0 {
		log.Printf("holding has none, check the input")
		return nil
	}

	for _, h := range holdings {
		if err := h.Run(); err != nil {
			log.Printf("run %s failed", h.Symbol)
			continue
		}
	}

	return nil
}

// Run launches all runners
func (h *Holding) Run() error {
	if h.Symbol == "" {
		log.Printf("symbol is nil")
		return nil
	}

	if len(h.Runners) == 0 {
		log.Printf("%s has no runner, skip", h.Symbol)
		return nil
	}

	for _, runner := range h.Runners {
		runner.Symbol = h.Symbol

		m, err := runner.Run()
		if err != nil {
			log.Printf("run %s for runner %s failed: %v",
				runner.Symbol, runner.Name, err)
			continue
		}

		if m.IsShouldRemind {
			// tempz
			fmt.Printf("%s: %s", h.Symbol, m.Message)
		}
	}

	return nil
}
