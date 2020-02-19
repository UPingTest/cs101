package main

import "testing"

func TestFizzbuzz(t *testing.T) {
	tests := []struct {
		a int
		r string
	}{
		{a: 0, r: ""},
		{a: 1, r: ""},
		{a: 3, r: "Fizz"},
		{a: 5, r: "Buzz"},
		{a: 15, r: "FizzBuzz"},
	}

	for row, test := range tests {
		r := Fizzbuzz(test.a)
		if r != test.r {
			t.Errorf("r is expected to be %s but now %s, row: %d", test.r, r, row)
		}
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		a int
		r bool
	}{
		{a: 0, r: false},
		{a: 1, r: false},
		{a: 2, r: true},
		{a: 3, r: true},
		{a: 4, r: false},
		{a: 97, r: true},
	}

	for row, test := range tests {
		r := IsPrime(test.a)
		if r != test.r {
			t.Errorf("r is expected to be %t but not %t, row: %d", test.r, r, row)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		a string
		r bool
	}{
		{a: "The quick brown 狐 jumped over the lazy 犬", r: false},
		{a: "aha", r: true},
		{a: "牛奶奶牛", r: true},
	}

	for row, test := range tests {
		r := IsPalindrome(test.a)
		if r != test.r {
			t.Errorf("r is expected to be %t but now %t, row: %d", test.r, r, row)
		}
	}
}
