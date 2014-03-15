package euclid

import (
	"testing"
	"testing/quick"
)

var config *quick.Config = &quick.Config{MaxCount: 1000}

func TestGcdOfOne(t *testing.T) {
	f := func(q uint) bool {
		return gcd(1, q) == 1
	}
	if err := quick.Check(f, config); err != nil {
		t.Error(err)
	}
}

func TestGcdOfZero(t *testing.T) {
	f := func(q uint) bool {
		return gcd(0, q) == q
	}
	if err := quick.Check(f, config); err != nil {
		t.Error(err)
	}
}

func TestCommutativity(t *testing.T) {
	f := func(p, q uint) bool {
		return gcd(p, q) == gcd(q, p)
	}
	if err := quick.Check(f, config); err != nil {
		t.Error(err)
	}
}

func TestGcdIsLessOrEqual(t *testing.T) {
	f := func(p, q uint) bool {
		g := gcd(p, q)
		return g <= p && g <= q
	}
	if err := quick.Check(f, config); err != nil {
		t.Error(err)
	}
}

func TestGcd(t *testing.T) {
	if gcd(2, 4) != 2 {
		t.Error("gcd(2, 4) != 2")
	}
}
