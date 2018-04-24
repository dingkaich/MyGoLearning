package mylogrus

import (
	"testing"
	"time"
)

func TestMylogrusMain1(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MylogrusMain1()
		})
	}
}

func BenchmarkMylogrusMain2(t *testing.B) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			MylogrusMain1()
		})
	}
}

func Benchmark_fmt(t *testing.B) {
	for i := 0; i < t.N; i++ { //use b.N for looping
		time.Sleep(time.Millisecond * 10)

	}

}
