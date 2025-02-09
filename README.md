# How I arrived at solution
1. Naive solution
2. Caching by maintaining state in EratosthenesSieve struct
3. Increasing blockSize * 2 after processing a block
4. Increasing blockSize as much as possible after processing a block
5. Honing in on optimal maxBlockSize
6. Preallocating plenty of space for primes slice

# Further Possible Optimizations
1. Cache line alignment of slices
