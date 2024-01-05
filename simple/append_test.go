package simple

import (
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	type args struct {
		i []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 01",
			args: args{
				i: []int{1, 2, 3, 4, 5, 6, 7, 8},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Append(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Append() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendRev(t *testing.T) {
	type args struct {
		i []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 01",
			args: args{
				i: []int{1, 2, 3, 4, 5, 6, 7, 8},
			},
			want: []int{8, 7, 6, 5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendRev(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendRev() = %v, want %v", got, tt.want)
			}
		})
	}
}
