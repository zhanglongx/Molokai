// Copyright 2020 Longxiao Zhang <zhanglongx@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a GPLv3-style
// license that can be found in the LICENSE file.

package core

type Symbol string

type Holding struct {
	Symbol Symbol
	Cost   float64

	Runner interface{}
}
