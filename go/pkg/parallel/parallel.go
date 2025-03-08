package parallel

import (
	"runtime"
	"sync"
)

func GetNumThreads(workUnits int) int {
	numThreads := runtime.NumCPU()

	if workUnits < numThreads {
		numThreads = workUnits
	}

	return numThreads
}

func For(workUnits int, loopBody func(i, threadId int)) {
	numThreads := GetNumThreads(workUnits)

	var wg sync.WaitGroup
	wg.Add(numThreads)

	for j := 0; j < numThreads; j++ {
		go func(threadId int) {
			defer wg.Done()

			batchSize := workUnits / numThreads
			start := batchSize * threadId
			stop := batchSize * (threadId + 1)

			if threadId == numThreads-1 {
				stop += workUnits % numThreads
			}

			for k := start; k < stop; k++ {
				loopBody(k, threadId)
			}

		}(j)
	}

	wg.Wait()
}
