// Copyright 2020 Longxiao Zhang <zhanglongx@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a GPLv3-style
// license that can be found in the LICENSE file.

package runner

import (
	"testing"

	"github.com/fxtlabs/date"
	"github.com/zhanglongx/Molokai/common"
)

func TestReminderRun(t *testing.T) {
	type args struct {
		symbol common.Symbol
		date   date.Date
		cost   float64
		params string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Bad JSON", args: args{
				symbol: "000001.SZ", date: date.Today(), cost: 1.1, params: "{\"Name\": }}",
			}, wantErr: true,
		},
		{
			name: "Bad Params", args: args{
				symbol: "000001.SZ", date: date.Today(), cost: 1.1, params: "{\"Name\": \"Percent\", \"Params\": {\"Maxttt\": 25}}",
			}, wantErr: true,
		},
		{
			name: "Good Params", args: args{
				symbol: "000001.SZ", date: date.Today(), cost: 1.1, params: "{\"Name\": \"Percent\", \"Params\": {\"Max\": 25}}",
			}, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ReminderRun(tt.args.symbol, tt.args.date, tt.args.cost, &tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReminderRun() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
