package graph

import "testing"

func TestFindLargestNumberOfNodes(t *testing.T) {
	type args struct {
		graph map[string][]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 01",
			args: args{
				graph: map[string][]string{
					"0": {"8", "1", "5"},
					"1": {"0"},
					"2": {"3", "4"},
					"3": {"2", "4"},
					"4": {"3", "2"},
					"5": {"0", "8"},
					"8": {"0", "5"},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLargestNumberOfNodes(tt.args.graph); got != tt.want {
				t.Errorf("FindLargestNumberOfNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
