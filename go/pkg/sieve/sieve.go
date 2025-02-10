package sieve

type Sieve interface {
	// A function that returns the nth 0-indexed prime number,
	// where 2 is the first prime number
	NthPrime(n int64) int64
}

// Note: this is the maximum block size and heavily impacts performance.
// May require dialing in based on specifically your hardware, eg. memory, CPU cache sizes
const maxBlockSize = 1 << 24

// A prime number Sieve implementation using a segmented Eratosthenes algorithm
type eratosthenesSieve struct {
	primes       []int64
	isNotPrime   []bool // Marks primes for the current Block
	blockStart   int64
	maxBlockSize int64
}

// A function that produces the nth 0-indexed prime number,
// using a segmented Eratosthenes algorithm.
// Previously computed primes are cached in memory.
// Numbers less than 0 will return the first prime number, 2.
func (eraSieve *eratosthenesSieve) NthPrime(n int64) int64 {
	if n < 0 {
		return 2
	}

	if n < int64(len(eraSieve.primes)) {
		return eraSieve.primes[n]
	}

	for int64(len(eraSieve.primes))-1 < n {

		var blockSize int64

		squared := eraSieve.blockStart * eraSieve.blockStart

		if squared <= eraSieve.blockStart+eraSieve.maxBlockSize {
			blockSize = squared
		} else {
			blockSize = eraSieve.maxBlockSize
		}

		eraSieve.markNonPrimes(blockSize)

		eraSieve.addPrimes(blockSize)

		eraSieve.blockStart += blockSize
	}

	return eraSieve.primes[n]
}

// adds any marked prime numbers to the sieve
// and resets indices to the default of false
func (eraSieve *eratosthenesSieve) addPrimes(blockSize int64) {
	for i := int64(0); i < blockSize; i++ {
		if !eraSieve.isNotPrime[i] {
			number := i + eraSieve.blockStart
			eraSieve.primes = append(eraSieve.primes, number)
		} else {
			eraSieve.isNotPrime[i] = false
		}
	}
}

// marks non-prime numbers in the current block of the sieve
func (eraSieve *eratosthenesSieve) markNonPrimes(blockSize int64) {

	// Marking all multiples of primes within the block as not prime
	for _, prime := range eraSieve.primes {

		multiplier := eraSieve.blockStart / prime
		if multiplier*prime < eraSieve.blockStart {
			multiplier++
		}

		// prime * multiplier == multiple
		offset := prime*multiplier - eraSieve.blockStart

		for offset < blockSize {
			eraSieve.isNotPrime[offset] = true
			offset += prime
		}
	}
}

var precalculatedPrimes = []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

func NewSieve() Sieve {
	eraSieve := eratosthenesSieve{}
	eraSieve.primes = make([]int64, 10, 100_000_000+1)
	eraSieve.isNotPrime = make([]bool, maxBlockSize)

	// seeding the EratosthenesSieve with precalculated values
	// so segmenting can be used straight away
	copy(eraSieve.primes, precalculatedPrimes[:])

	// start algorithm at the number after the last known prime
	eraSieve.blockStart = precalculatedPrimes[len(precalculatedPrimes)-1] + 1
	eraSieve.maxBlockSize = maxBlockSize

	return &eraSieve
}
