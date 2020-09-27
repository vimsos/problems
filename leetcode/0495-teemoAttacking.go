package leetcode

// In LOL world, there is a hero called Teemo and his attacking can make his enemy Ashe be in poisoned condition. Now, given the Teemo's attacking ascending time series towards Ashe and the poisoning time duration per Teemo's attacking, you need to output the total time that Ashe is in poisoned condition.
// You may assume that Teemo attacks at the very beginning of a specific time point, and makes Ashe be in poisoned condition immediately.

func findPoisonedDuration(timeSeries []int, duration int) int {
	l := len(timeSeries)
	if l == 0 {
		return 0
	}

	total := duration
	for i := 0; i < l-1; i++ {
		dif := timeSeries[i+1] - timeSeries[i]
		if dif < duration {
			total += dif
		} else {
			total += duration
		}
	}
	return total
}
