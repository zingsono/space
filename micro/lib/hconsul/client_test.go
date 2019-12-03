package hconsul

import (
	"testing"
)

func TestRegister(t *testing.T) {
	type args struct {
		datacenter string
		node       string
		address    string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Register", args: args{datacenter: "", node: "ms", address: "192.168.0.2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Register(tt.args.datacenter, tt.args.node, tt.args.address)
		})
	}
}
