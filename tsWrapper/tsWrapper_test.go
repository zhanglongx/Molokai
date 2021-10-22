package tsWrapper

import (
	"testing"

	"github.com/fxtlabs/date"
)

func TestTsWrapper_AdjFactor(t *testing.T) {
	type fields struct {
		Token string
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

			Token = tt.fields.Token

			_, err := AdjFactor(tt.args.tsCode, tt.args.date)
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
			name: "Bad", wantErr: true,
			fields: fields{Token: "f4673f7862e73483c5e65cd9a036eedd39e72d484194a85dabcf958b"},
			args:   args{tsCode: "100001.SZ"},
		},
		{
			name: "Good", wantErr: false,
			fields: fields{Token: "f4673f7862e73483c5e65cd9a036eedd39e72d484194a85dabcf958b"},
			args:   args{tsCode: "000001.SZ"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			Token = tt.fields.Token

			_, err := RecentClose(tt.args.tsCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("TsWrapper.RecentClose() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSymbol(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Bad", args: args{s: "100001.SZ"}, want: "", wantErr: true},
		{name: "Good1", args: args{s: "000001.SZ"}, want: "000001.SZ", wantErr: false},
		{name: "Good2", args: args{s: "000001"}, want: "000001.SZ", wantErr: false},
		{name: "Good3", args: args{s: "格力电器"}, want: "000651.SZ", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Token = "f4673f7862e73483c5e65cd9a036eedd39e72d484194a85dabcf958b"

			got, err := Symbol(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Symbol() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Symbol() = %v, want %v", got, tt.want)
			}
		})
	}
}
