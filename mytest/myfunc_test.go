package mytest

import "testing"

func TestDivision(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "1", args: args{a: 21, b: 123}, want: 0, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Division(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Division() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Division() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Division((float64)(i+1), (float64)(i))
	}
}
