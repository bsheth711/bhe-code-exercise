package parallel

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNumThreadsLarge(t *testing.T) {
	numCPU := runtime.NumCPU()
	workUnits := 10000000

	if numCPU < workUnits {
		assert.Equal(t, numCPU, GetNumThreads(workUnits))
	} else {
		assert.Equal(t, workUnits, GetNumThreads(workUnits))
	}
}

func TestGetNumThreadsSmall(t *testing.T) {
	numCPU := runtime.NumCPU()
	workUnits := 1

	if numCPU < workUnits {
		assert.Equal(t, numCPU, GetNumThreads(workUnits))
	} else {
		assert.Equal(t, workUnits, GetNumThreads(workUnits))
	}
}

func TestFor(t *testing.T) {

	n := 100000

	total1 := 0

	for i := 0; i < n; i++ {
		total1 += i
	}

	total2 := 0
	totals := make([]int, GetNumThreads(n))

	For(n, func(i, threadId int) {
		totals[threadId] += i
	})

	for _, total := range totals {
		total2 += total
	}

	assert.Equal(t, total1, total2)
}
