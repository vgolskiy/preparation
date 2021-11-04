package main

import "fmt"

func tournamentWinner(competitions *[][2]string, results *[]int) string {
	var winner string
	winners := make(map[string]int)

	for i, teams := range *competitions {
		team := teams[1-(*results)[i]]
		if _, found := winners[team]; found {
			winners[team] += 3
		} else {
			winners[team] = 3
		}
		if winner != "" {
			if winners[team] > winners[winner] {
				winner = team
			}
		} else {
			winner = team
		}
	}
	return winner
}

var competitions = [][2]string{{"HTML", "C#"}, {"C#", "Python"}, {"Python", "HTML"}}
var results = []int{0, 0, 1}

func main() {
	fmt.Println(tournamentWinner(&competitions, &results))
}
