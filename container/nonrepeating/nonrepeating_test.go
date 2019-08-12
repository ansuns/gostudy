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
			t.Errorf("got %d for input %s; " +
				"expect %d", actual, tt.s, tt.ans)
		}
	}
}
