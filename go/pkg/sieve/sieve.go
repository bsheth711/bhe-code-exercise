package sieve

type Sieve interface {
	NthPrime(n int64) int64
}

type EratosthenesSieve struct {
	primes       []int64
	blockStart   int64
	blockSize    int64
	maxBlockSize int64
}

func (eraSieve EratosthenesSieve) NthPrime(n int64) int64 {
	if n < 0 {
		return 2
	}

	if n < int64(len(eraSieve.primes)) {
		return eraSieve.primes[n]
	}

	for int64(len(eraSieve.primes))-1 < n {

		blockEnd := eraSieve.blockStart + eraSieve.blockSize
		var isNotPrime []bool = make([]bool, eraSieve.blockSize)

		for _, prime := range eraSieve.primes {

			multiplier := eraSieve.blockStart / prime

			multiple := prime * multiplier

			if multiple < eraSieve.blockStart {
				multiplier++
				multiple = prime * multiplier
			}

			for multiple < blockEnd {
				offset := multiple - eraSieve.blockStart

				isNotPrime[offset] = true

				multiplier++
				multiple = prime * multiplier
			}
		}

		for i := range isNotPrime {
			if !isNotPrime[i] {
				number := int64(i) + eraSieve.blockStart
				eraSieve.primes = append(eraSieve.primes, number)
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

	return eraSieve.primes[n]
}

func NewSieve() Sieve {
	eraSieve := EratosthenesSieve{}
	eraSieve.primes = make([]int64, 10, 200000001)

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

	// starting block size, will expand to maxBlockSize
	eraSieve.blockSize = 1 << 9

	// Note: this is the maximum block size and heavily impacts performance.
	// May require dialing in based on specifically your hardware, eg. memory, CPU cache sizes
	// Currently, it is set to the optimal value for my M1 Macbook Air: 100000000 take ~12s
	eraSieve.maxBlockSize = 1 << 24

	return eraSieve
}
