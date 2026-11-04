package hello

import "testing"

func TestGreet(t *testing.T) {
	for _, tc := range []struct{ name, want string }{
		{"小林", "Hello, 小林!"},
		{"", "Hello, world!"},
	} {
		if got := Greet(tc.name); got != tc.want {
			t.Errorf("Greet(%q) = %q, want %q", tc.name, got, tc.want)
		}
	}
}
