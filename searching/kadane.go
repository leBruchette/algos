package searching

/*
   local_maximum =  max(A[i], A[i] + local_maximum)
*/

func KadaneSearch(numbers []int) int {
	//log.SetOutput(io.Discard)
	if len(numbers) == 0 {
		return 0
	}

	globalMax, localMax := 0, 0
	if numbers[0] > 0 {
		globalMax = numbers[0]
		localMax = numbers[0]
	}

	for i := 1; i < len(numbers); i++ {
		localMax = max(numbers[i], numbers[i]+localMax)
		globalMax = max(localMax, globalMax)
	}

	return globalMax

}
