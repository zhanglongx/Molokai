package tsWrapper

import (
	"testing"

	"github.com/ShawnRong/tushare-go"
	"github.com/fxtlabs/date"
)

func TestTsWrapper_AdjFactor(t *testing.T) {
	type fields struct {
		Token string
		api   *tushare.TuShare
	}
	type args struct {
		tsCode string
		date   date.Date
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Bad Token", wantErr: true,
			fields: fields{Token: "f4673f7862e73483c5e65cd9a036eedd39e72d484194a85dabcf958c"},
			args:   args{tsCode: "000001.SZ", date: date.Today()},
		},
		{
			name: "Bad Symbol", wantErr: true,
			fields: fields{Token: "f4673f7862e73483c5e65cd9a036eedd39e72d484194a85dabcf958b"},
			args:   args{tsCode: "1000001.SZ", date: date.Today()},
		},
		{
			name: "Bad Date", wantErr: true,
			fields: fields{Token: "f4673f7862e73483c5e65cd9a036eedd39e72d484194a85dabcf958b"},
			args:   args{tsCode: "000001.SZ", date: date.New(2021, 8, 15)},
		},
		{
			name: "Good Symbol", wantErr: false,
			fields: fields{Token: "f4673f7862e73483c5e65cd9a036eedd39e72d484194a85dabcf958b"},
			args:   args{tsCode: "000002.SZ", date: date.New(2021, 1, 13)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TsWrapper{
				Token: tt.fields.Token,
				api:   tt.fields.api,
			}

			ts.Init()

			_, err := ts.AdjFactor(tt.args.tsCode, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("TsWrapper.AdjFactor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTsWrapper_RecentClose(t *testing.T) {
	type fields struct {
		Token string
		api   *tushare.TuShare
	}
	type args struct {
		tsCode string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Good", wantErr: true,
			fields: fields{Token: "f4673f7862e73483c5e65cd9a036eedd39e72d484194a85dabcf958b"},
			args:   args{tsCode: "000001.SZ"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TsWrapper{
				Token: tt.fields.Token,
				api:   tt.fields.api,
			}

			ts.Init()

			_, err := ts.RecentClose(tt.args.tsCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("TsWrapper.RecentClose() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
