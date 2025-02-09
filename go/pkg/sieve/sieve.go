package sieve

type Sieve interface {
	NthPrime(n int64) int64
}

const startingBlockSize = 1 << 9

// Note: this is the maximum block size and heavily impacts performance.
// May require dialing in based on specifically your hardware, eg. memory, CPU cache sizes
// Currently, it is set to the optimal value for my M1 Macbook Air: 100000000 take ~12s
const maxBlockSize = 1 << 24

type eratosthenesSieve struct {
	Primes       []int64
	isNotPrime   []bool
	blockStart   int64
	blockSize    int64
	maxBlockSize int64
}

func (eraSieve eratosthenesSieve) NthPrime(n int64) int64 {
	if n < 0 {
		return 2
	}

	if n < int64(len(eraSieve.Primes)) {
		return eraSieve.Primes[n]
	}

	for int64(len(eraSieve.Primes))-1 < n {

		blockEnd := eraSieve.blockStart + eraSieve.blockSize

		for _, prime := range eraSieve.Primes {

			multiplier := eraSieve.blockStart / prime

			multiple := prime * multiplier

			if multiple < eraSieve.blockStart {
				multiplier++
				multiple = prime * multiplier
			}

			for multiple < blockEnd {
				offset := multiple - eraSieve.blockStart

				eraSieve.isNotPrime[offset] = true

				multiplier++
				multiple = prime * multiplier
			}
		}

		for i := int64(0); i < eraSieve.blockSize; i++ {
			if !eraSieve.isNotPrime[i] {
				number := i + eraSieve.blockStart
				eraSieve.Primes = append(eraSieve.Primes, number)
			} else {
				eraSieve.isNotPrime[i] = false
			}
		}

		eraSieve.blockStart += eraSieve.blockSize

		nextBlockSize := eraSieve.blockSize * 2

		for (eraSieve.blockSize < eraSieve.maxBlockSize) &&
			(eraSieve.blockStart*eraSieve.blockStart > eraSieve.blockStart+nextBlockSize) {
			eraSieve.blockSize = nextBlockSize
			nextBlockSize <<= 1
		}
	}

	return eraSieve.Primes[n]
}

func NewSieve() Sieve {
	eraSieve := eratosthenesSieve{}
	eraSieve.Primes = make([]int64, 10, 100000001)
	eraSieve.isNotPrime = make([]bool, maxBlockSize)

	// seeding the EratosthenesSieve with precalculated values
	// so segmenting can be used straight away
	eraSieve.Primes[0] = 2
	eraSieve.Primes[1] = 3
	eraSieve.Primes[2] = 5
	eraSieve.Primes[3] = 7
	eraSieve.Primes[4] = 11
	eraSieve.Primes[5] = 13
	eraSieve.Primes[6] = 17
	eraSieve.Primes[7] = 19
	eraSieve.Primes[8] = 23
	eraSieve.Primes[9] = 29

	// next number to start algorithm at after the last known prime
	eraSieve.blockStart = 30

	// starting block size, will expand to maxBlockSize
	eraSieve.blockSize = startingBlockSize
	eraSieve.maxBlockSize = maxBlockSize

	return eraSieve
}
