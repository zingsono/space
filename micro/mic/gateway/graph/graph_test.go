package graph

import (
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"
)

func TestArgument(t *testing.T) {
	type args struct {
		args []graphql.FieldConfigArgument
	}
	tests := []struct {
		name string
		args args
		want graphql.FieldConfigArgument
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Argument(tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Argument() = %v, want %v", got, tt.want)
			}
		})
	}
}
