package main

type diracDice struct {
	// roll times
	times int
	// current dice number
	c int
	// which player to roll
	// 0 for player 1, 1 for player 2
	playerTurn int
	// score of each player
	playerScore [2]int
	// which position of each player
	playerPos [2]int
}

func main() {
	results := make(map[diracDice][2]int)
	var state = diracDice{playerPos: [2]int{10, 2}}
	res := roll(state, results)
	if res[0] > res[1] {
		println(res[0])
	} else {
		println(res[1])
	}
}

func roll(state diracDice, results map[diracDice][2]int) (r [2]int) {
	// cache
	if res, ok := results[state]; ok {
		return res
	}
	defer func() { results[state] = r }()
	turn := state.playerTurn
	if state.times == 3 {
		state.playerPos[turn] = round(state.playerPos[turn] + state.c)
		state.playerScore[turn] += state.playerPos[turn]
		if state.playerScore[turn] >= 21 {
			r[turn] = 1
			return r
		}
		// reset state , turn to next player
		state.c = 0
		state.times = 0
		// Trick:change the turn : 0 -> 1, 1 -> 0
		state.playerTurn = 1 - state.playerTurn
		return roll(state, results)
	}
	for i := 1; i <= 3; i++ {
		state.times++
		state.c += i
		res := roll(state, results)
		r[0] += res[0]
		r[1] += res[1]
		state.times--
		state.c -= i
	}
	return r
}

func round(input int) int {
	if input%10 == 0 {
		return 10
	}
	return input % 10
}
