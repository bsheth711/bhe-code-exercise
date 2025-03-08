# Implementation Details
I implemented a segmented Eratosthenes Sieve. 
Primes are cached within the sieve to improve performance of runs within the same context.  

Note: Correctly setting maxBlockSize is extremely important: too large and you lose caching benefits 
and may run in to memory issues. Too small and you lose out on CPU optimizations.
Based on your hardware, you may need to change the value for optimal performance. 
Currently it is set to the optimal value for my M1 Macbook Air 
(getting the 100_000_000th prime takes ~4.1s).   

Note: In multiple places I square one side of an equation to avoid square rooting.
This runs the risk of integer overflow for very large values, but gives great performance gains.
As such, the program will only work for primes below ~3 billion.

# A Note On Multithreading
I implemented multithreading on the branch feature-bsheth-multithreading. 
The results of multithreading are variable. On my Windows Desktop, multithreading has a beneficial impact (~200-500ms). 
However on my M1 Macbook Air, multithreading has no impact or a slightly negative impact. 
This is likely due differences in compilation and architecture between the two systems. 
Of course when thinking about optimization and multithreading, it is important to think about your target system 
and test, test, test to ensure you are getting the results you expect. 
In any case, the algorithm I am using is already likely to close to memory bounded--ie. limited by the speed of memory accesses.

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
