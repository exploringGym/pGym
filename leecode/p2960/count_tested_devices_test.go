package p2960

import "testing"

func Test_countTestedDevices(t *testing.T) {
	type args struct {
		batteryPercentages []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 01",
			args: args{
				batteryPercentages: []int{1, 1, 2, 1, 3},
			},
			want: 3,
		},
		{
			name: "Test 02",
			args: args{
				batteryPercentages: []int{0, 1, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTestedDevices(tt.args.batteryPercentages); got != tt.want {
				t.Errorf("countTestedDevices() = %v, want %v", got, tt.want)
			}
		})
	}
}
