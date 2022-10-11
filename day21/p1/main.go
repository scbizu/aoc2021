package main

type diceCounter struct {
	c     int
	times int
}

func (dc *diceCounter) roll() int {
	var res int
	for i := 0; i < 3; i++ {
		res += dc.c
		dc.c++
		dc.times++
	}
	return res
}

func round(input int) int {
	if input%10 == 0 {
		return 10
	}
	return input % 10
}

func main() {
	p1, p2 := 10, 2
	var s1, s2 int
	var loser int
	var endTime int
	const scoreLimit = 1000
	dice := &diceCounter{c: 1}
	for {
		if s2 < scoreLimit {
			r1 := dice.roll()
			p1 = round(p1 + r1)
			s1 += p1
			println("Player 1: ", p1, "score: ", s1)
		} else {
			loser = s1
			endTime = dice.times
			break
		}
		if s1 < scoreLimit {
			r2 := dice.roll()
			p2 = round(r2 + p2)
			s2 += p2
			println("Player 2: ", p2, "score: ", s2)
		} else {
			loser = s2
			endTime = dice.times
			break
		}
	}
	println(loser * endTime)
}
