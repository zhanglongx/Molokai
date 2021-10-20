// Copyright 2020 Longxiao Zhang <zhanglongx@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a GPLv3-style
// license that can be found in the LICENSE file.

package core

import (
	"log"

	"github.com/zhanglongx/Molokai/runner"
)

type Holding struct {
	Symbol string

	Runners []runner.Runner
}

func RunHoldings(holdings []Holding) error {
	for _, h := range holdings {
		if err := h.Run(); err != nil {
			log.Printf("run %s failed", h.Symbol)
			continue
		}
	}

	return nil
}

func (h *Holding) Run() error {
	for _, runner := range h.Runners {
		runner.Symbol = h.Symbol

		if _, err := runner.Run(); err != nil {
			log.Printf("run %s for runner %s failed: %v",
				runner.Symbol, runner.Name, err)
		}
	}

	return nil
}
