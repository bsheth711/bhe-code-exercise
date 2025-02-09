package sieve

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthPrimeSmall(t *testing.T) {
	sieve := NewSieve()

	assert.Equal(t, int64(2), sieve.NthPrime(0))
	assert.Equal(t, int64(3), sieve.NthPrime(1))
	assert.Equal(t, int64(5), sieve.NthPrime(2))
	assert.Equal(t, int64(71), sieve.NthPrime(19))
	assert.Equal(t, int64(2909), sieve.NthPrime(420))
	assert.Equal(t, int64(3581), sieve.NthPrime(500))
	assert.Equal(t, int64(2017), sieve.NthPrime(305))
	assert.Equal(t, int64(17393), sieve.NthPrime(2000))
}

func TestNthPrime(t *testing.T) {
	sieve := NewSieve()

	assert.Equal(t, int64(2), sieve.NthPrime(0))
	assert.Equal(t, int64(71), sieve.NthPrime(19))
	assert.Equal(t, int64(541), sieve.NthPrime(99))
	assert.Equal(t, int64(3581), sieve.NthPrime(500))
	assert.Equal(t, int64(7793), sieve.NthPrime(986))
	assert.Equal(t, int64(17393), sieve.NthPrime(2000))
	assert.Equal(t, int64(15485867), sieve.NthPrime(1_000_000))
	assert.Equal(t, int64(179424691), sieve.NthPrime(10_000_000))
	assert.Equal(t, int64(2038074751), sieve.NthPrime(100_000_000)) //not required, just a fun challenge
	// I love challenges ;)
}

func FuzzNthPrime(f *testing.F) {
	sieve := NewSieve()

	f.Fuzz(func(t *testing.T, n int64) {
		if !big.NewInt(sieve.NthPrime(n)).ProbablyPrime(0) {
			t.Errorf("the sieve produced a non-prime number at index %d", n)
		}
	})
}
