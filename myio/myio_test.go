package myio

import "testing"

func TestMyioMain(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"dingkai"},
	}
	for range tests {
		MyioMain()
	}
}

func Testaa(t *testing.T) {
	i := channelLens(20, 12, 12)
	t.Errorf("%d", i)
	i = channelLens(10, 12, 12)
	t.Errorf("%d", i)

}
