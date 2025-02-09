# Implementation Details
I implemented a segmented Eratosthenes Sieve. 
Primes are cached within the sieve to improve performace of runs within the same context.
Note: Correctly setting maxBlockSize is extremely important: too large and you lose caching benefits 
and may run in to memory issues. Too small and you branch too often, losing out on CPU optimizations.
Based on your hardware, you may need to change the value for optimal performance. 
Currently it is set to the optimal value for my M1 Macbook Air. 

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

# Other attempted optimizations
1. Memory aligning slices - little to no effect

# Additional possible optimizations
1. Inspection of assembly to ensure SIMD usage
