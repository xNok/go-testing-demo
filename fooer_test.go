package main

import "testing"

func TestFooer(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fooer(tt.args.input); got != tt.want {
				t.Errorf("Fooer() = %v, want %v", got, tt.want)
			}
		})
	}
}
