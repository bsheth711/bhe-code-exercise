# Implementation Details
I implemented a segmented Eratosthenes Sieve. 
Primes are cached within the sieve to improve performace of runs within the same context.  

Note: Correctly setting maxBlockSize is extremely important: too large and you lose caching benefits 
and may run in to memory issues. Too small and you lose out on CPU optimizations.
Based on your hardware, you may need to change the value for optimal performance. 
Currently it is set to the optimal value for my M1 Macbook Air 
(getting the 100_000_000th prime takes ~5s).   

Note: In multiple places I square one side of an equation to avoid square rooting.
This runs the risk of integer overflow for very large values, but gives great performance gains.
As such, the program will only work for primes below ~3 billion.

# How To Run
1. From the repository folder move to the go folder
```
cd ./go
```
2. Run tests
```
go test ./...
```

# How I arrived at the implementation
1. Naive solution
2. Caching by maintaining state in EratosthenesSieve struct
3. Increasing blockSize * 2 after processing a block - very large effect
4. Increasing blockSize as much as possible after processing a block - very large effect
5. Honing in on optimal maxBlockSize - very large effect
6. Preallocating plenty of space for primes slice - small effect
7. Preallocating space for isNotPrime - little to no effect
8. Only checking primes up to sqrt(blockEnd) - very large effect

# Other attempted optimizations
1. Memory aligning slices - little to no effect

# Additional possible optimizations
- Inspection of assembly to ensure SIMD usage
- Usage of wheel algorithm
