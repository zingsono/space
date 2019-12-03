package hconsul

import (
	"testing"
)

func TestReg(t *testing.T) {
	type args struct {
		host string
		ip   string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "intention", args: args{host: "", ip: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Reg(tt.args.host, tt.args.ip)
		})
	}
}
