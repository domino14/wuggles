package main

func boardBonus(idx int, round int) int {
	if round == 1 {
		return 1 // no bonuses
	}
	// Find the multiplier bonus for this square index.
	x, y := indexToXY(idx, roundToDim(round))
	switch round {
	case 2:
		if (x == 2 && y == 0) || (x == 3 && y == 5) {
			return 2
		}
		return 1
	case 3:
		if (x == 0 && y == 0) || (x == 5 && y == 5) {
			return 3
		}
		return 1
	case 4:
		if (x == 0 && y == 0) || (x == 5 && y == 5) {
			return 2
		}
		if (x == 5 && y == 0) || (x == 0 && y == 5) {
			return 3
		}
		return 1
	}
	return 1 // probably not right
}

func score(word []rune, multiplier int) int {
	wl := len(word)
	scoreMap := map[int]int{
		3: 1, 4: 2, 5: 4, 6: 7, 7: 11, 8: 15, 9: 20, 10: 25,
		11: 30, 12: 36, 13: 42, 14: 48, 15: 60,
	}
	return scoreMap[wl] * multiplier
}
