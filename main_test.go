package main

import "testing"

type Test struct {
	in    []byte
	out   int
	haser bool
}

var tests = []Test{
	{[]byte("lol"), 3, true},
	{[]byte("t"), 1, true},
	{[]byte{0xff}, 0, false},
}

func TestGetUTFLength(t *testing.T) {
	for i, test := range tests {
		size, err := GetUTFLength(test.in)
		if (err != nil) == test.haser {
			t.Errorf("expected error: , got: %v", err != nil)
		}
		if size != test.out {
			t.Errorf("#%d: Size(%s)=%d; want %d", i, test.in, size, test.out)
		}
	}
}
