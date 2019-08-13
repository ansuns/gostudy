package main

import "testing"

func TestNonrepeating(t *testing.T) {
	test := []struct {
		s   string //分析的字符串
		ans int    //答案
	}{
		{"abcabcbb", 3},
		{"pwwkew", 3},
		{"", 0},
		{"123123", 3},

		//中文
		{"我爱中国", 4},
		{"一二三一二三", 3},

		//特殊
		{".{·】", 4},
	}

	for _, tt := range test {
		actual := length0fNonRepeatingSubStr(tt.s);
		if actual != tt.ans {
			t.Errorf("got %d for input %s; "+
				"expect %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花，"
	ans := 8
	b.ResetTimer() //准备数据的开始计算时间重置
	for i := 0; i < b.N; i++ {
		actual := length0fNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; "+
				"expect %d", actual, s, ans)
		}
	}
}
