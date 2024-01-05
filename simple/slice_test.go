package simple

import (
	"reflect"
	"testing"
)

func TestSliceFront(t *testing.T) {
	type args struct {
		x   []int
		cut int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 01",
			args: args{
				x:   []int{1, 2, 3, 4, 5, 6, 7, 8},
				cut: 2,
			},
			want: []int{3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceFront(tt.args.x, tt.args.cut); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceEnd(t *testing.T) {
	type args struct {
		x   []int
		cut int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 01",
			args: args{
				x:   []int{1, 2, 3, 4, 5, 6, 7, 8},
				cut: 2,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceEnd(tt.args.x, tt.args.cut); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
