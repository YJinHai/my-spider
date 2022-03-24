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
		want float32
	}{
		{"Cosine sim 1", args{"<a target=\"_top\" class=\"main\" href=\"/monthly/2022/02\">\n              数据库内核月报 － 2022/02\n            </a>", "<a target=\"_top\" class=\"main\" href=\"/monthly/2022/01\">\n              数据库内核月报 － 2022/01\n            </a>"}, 0.0},
		{"Cosine sim 2", args{"I love horror movies", "Lights out is a horror movie"}, 0.20412414},
		{"Cosine sim 3", args{"love horror movies", "Lights out horror movie"}, 0.28867513},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CosineSimilarity(tt.args.str1, tt.args.str2, 0); got != tt.want {
				t.Errorf("CosineSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}
