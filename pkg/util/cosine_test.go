package util

import "testing"

func TestCosineSimilarity(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Cosine sim 1", args{"<a target=\"_top\" class=\"main\" href=\"/monthly/2022/02\">", "<a target=\"_top\" class=\"main\" href=\"/monthly/2020/02\">"}, 0.0},
		{"Cosine sim 2", args{"中", "国"}, 0.20412414},
		{"Cosine sim 3", args{"love horror movies", "Lights out horror movie"}, 0.28867513},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CosineSimilarity(tt.args.str1, tt.args.str2, 1); got != tt.want {
				t.Errorf("CosineSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}
