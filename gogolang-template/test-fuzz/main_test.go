package main

import (
	"strconv"
	"testing"
)

// 普通测试
func TestEqual(t *testing.T) {
	tests := []struct {
		name   string
		inputA []byte
		inputB []byte
		want   bool
	}{
		{"right case", []byte{'f', 'u', 'z', 'z'}, []byte{'f', 'u', 'z', 'z'}, true},
		{"right case", []byte{'a', 'b', 'c'}, []byte{'b', 'c', 'd'}, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.inputA, tt.inputB); got != tt.want {
				t.Error("expected " + strconv.FormatBool(tt.want) + ", got " + strconv.FormatBool(got))
			}
		})
	}
}

// go test -fuzz .
// go test -fuzz=Fuzz -fuzztime=10s .
func FuzzEqual(f *testing.F) {
	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		Equal(a, b)
	})
}
