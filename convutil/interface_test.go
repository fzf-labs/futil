package conv

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name      string
		args      args
		wantValue interface{}
		wantOk    bool
	}{
		{
			name:      "",
			args:      args{},
			wantValue: nil,
			wantOk:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotOk := Interface(tt.args.v)
			assert.Equalf(t, tt.wantValue, gotValue, "Interface(%v)", tt.args.v)
			assert.Equalf(t, tt.wantOk, gotOk, "Interface(%v)", tt.args.v)
		})
	}
}
