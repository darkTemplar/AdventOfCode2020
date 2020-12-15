package day15

func FindSpokenNumber(initialNumbers []int, finalTurn int) int {
	spoken := make(map[int][]int)
	lastSpoken := 0
	for i, n := range initialNumbers {
		lastSpoken = n
		speak(spoken, n, i+1)
	}
	// play the game till finalTurn
	for i := len(initialNumbers) + 1; i <= finalTurn; i++ {
		counts, _ := spoken[lastSpoken]
		if len(counts) == 1 {
			lastSpoken = 0
			speak(spoken, 0, i)
		} else {
			lastSpoken = counts[1] - counts[0]
			speak(spoken, counts[1]-counts[0], i)
		}
	}
	return lastSpoken
}

func speak(spoken map[int][]int, value int, position int) {
	if count, ok := spoken[value]; ok {
		count = append(count, position)
		// since we only keep track of last 2 turns
		if len(count) > 2 {
			count = count[1:]
		}
		spoken[value] = count
	} else {
		spoken[value] = []int{position}
	}
}
