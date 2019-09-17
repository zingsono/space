package hdate

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	t.Log(time.Now().Format(time.RFC3339))
}

func TestTime(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "Time",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Time(); got != tt.want {
				t.Errorf("Time() = %v, want %v", got, tt.want)
			}
		})
	}
}
