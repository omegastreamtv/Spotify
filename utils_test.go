package spotify

import (
	"testing"
)

func TestIsLeft(t *testing.T) {
	// prepare test cases
	tests := []struct {
		name   string
		either Either[int, string]
		want   bool
	}{
		{"Left value", Either[int, string]{isLeft: true, left: 10, right: ""}, true},
		{"Right value", Either[int, string]{isLeft: false, left: 0, right: "test"}, false},
	}

	// execute test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.either.IsLeft(); got != tt.want {
				t.Errorf("IsLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRight(t *testing.T) {
	tests := []struct {
		name   string
		either Either[int, string]
		want   bool
	}{
		{"Left value", Either[int, string]{isLeft: true, left: 10, right: ""}, false},
		{"Right value", Either[int, string]{isLeft: false, left: 0, right: "test"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.either.IsRight(); got != tt.want {
				t.Errorf("IsRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeft(t *testing.T) {
	tests := []struct {
		name   string
		either Either[int, string]
		want   int
		ok     bool
	}{
		{"Left value", Either[int, string]{isLeft: true, left: 10, right: ""}, 10, true},
		{"Right value", Either[int, string]{isLeft: false, left: 0, right: "test"}, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.either.Left()
			if got != tt.want || ok != tt.ok {
				t.Errorf("Left() = %v, %v, want %v, %v", got, ok, tt.want, tt.ok)
			}
		})
	}
}

func TestRight(t *testing.T) {
	tests := []struct {
		name   string
		either Either[int, string]
		want   string
		ok     bool
	}{
		{"Left value", Either[int, string]{isLeft: true, left: 10, right: ""}, "", false},
		{"Right value", Either[int, string]{isLeft: false, left: 0, right: "test"}, "test", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.either.Right()
			if got != tt.want || ok != tt.ok {
				t.Errorf("Right() = %v, %v, want %v, %v", got, ok, tt.want, tt.ok)
			}
		})
	}
}
