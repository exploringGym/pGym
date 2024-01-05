package recursion

import (
	"reflect"
	"testing"
)

func TestDecimalToBin(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Test 01",
			args: args{
				i: 2,
			},
			want: []byte{1, 0},
		},
		{
			name: "Test 02",
			args: args{
				i: 2000,
			},
			want: []byte{1, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecimalToBin(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecimalToBin() = %v, want %v", got, tt.want)
			}
		})
	}
}
