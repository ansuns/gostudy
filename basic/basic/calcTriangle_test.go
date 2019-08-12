package main

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
		{-310000, 40000, 50000},
	}

	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			//汇报错误的接口
			t.Errorf("caclTriangle (%d, %d); "+
				"got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}
