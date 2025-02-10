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

		if eraSieve.blockStart*eraSieve.blockStart <= eraSieve.blockStart+eraSieve.maxBlockSize {
			blockSize = eraSieve.blockStart * eraSieve.blockStart
		} else {
			blockSize = eraSieve.maxBlockSize
		}

		blockEnd := eraSieve.blockStart + blockSize

		// Marking all multiples of primes within the block as not prime
		for _, prime := range eraSieve.primes {

			multiplier := eraSieve.blockStart / prime
			if prime*multiplier < eraSieve.blockStart {
				multiplier++
			}

			multiple := prime * multiplier

			for multiple < blockEnd {
				offset := multiple - eraSieve.blockStart

				eraSieve.isNotPrime[offset] = true

				multiple += prime
			}
		}

		// Adding identified primes to sieve
		for i := int64(0); i < blockSize; i++ {
			if !eraSieve.isNotPrime[i] {
				number := i + eraSieve.blockStart
				eraSieve.primes = append(eraSieve.primes, number)
			} else {
				eraSieve.isNotPrime[i] = false
			}
		}

		eraSieve.blockStart += blockSize
	}

	return eraSieve.primes[n]
}

func NewSieve() Sieve {
	eraSieve := eratosthenesSieve{}
	eraSieve.primes = make([]int64, 10, 100_000_000+1)
	eraSieve.isNotPrime = make([]bool, maxBlockSize)

	// seeding the EratosthenesSieve with precalculated values
	// so segmenting can be used straight away
	eraSieve.primes[0] = 2
	eraSieve.primes[1] = 3
	eraSieve.primes[2] = 5
	eraSieve.primes[3] = 7
	eraSieve.primes[4] = 11
	eraSieve.primes[5] = 13
	eraSieve.primes[6] = 17
	eraSieve.primes[7] = 19
	eraSieve.primes[8] = 23
	eraSieve.primes[9] = 29

	// next number to start algorithm at after the last known prime
	eraSieve.blockStart = 30

	eraSieve.maxBlockSize = maxBlockSize

	return &eraSieve
}
