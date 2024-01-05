package graph

import "testing"

func TestCountGraphs(t *testing.T) {
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
					"1":  {"2"},
					"2":  {"1"},
					"3":  {},
					"4":  {"6"},
					"5":  {"6"},
					"6":  {"4", "5", "7", "8"},
					"7":  {"6"},
					"8":  {"6"},
					"9":  {},
					"10": {},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountGraphs(tt.args.graph); got != tt.want {
				t.Errorf("CountGraphs() = %v, want %v", got, tt.want)
			}
		})
	}
}
