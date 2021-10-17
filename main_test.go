package main

import "testing"

func Test_stringAdd(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "a1", args: args{a: "0.9", b: "0.9",}, want: "1.8"},
		{name: "a2", args: args{a: "0.1", b: "0.2",}, want: "0.3"},
		{name: "a3", args: args{a: "0.9999", b: "0.9999",}, want: "1.888"},
		{name: "a4", args: args{a: "0.9", b: "0.1",}, want: "1.0"},
		{name: "a5", args: args{a: "0.02", b: "0.092",}, want: "0"},
		{name: "a6", args: args{a: "0.99", b: "0.01",}, want: "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringAdd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("stringAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
