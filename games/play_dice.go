package games

import (
	"fmt"
	"math/rand"
)

type DiceGameModel struct {
	Name            int
	Score           int
	TotalDice       int
	Points          []int
	EvaluatedPoints []int
}

func PlayDice(totalPlayer, totalDice int) {
	var players []DiceGameModel
	gameRound := 1
	for {
		fmt.Println(fmt.Sprintf("Game ke %v", gameRound))
		//register player to a game
		if gameRound == 1 {
			for i := 0; i < totalPlayer; i++ {
				players = append(players, DiceGameModel{Name: i + 1, TotalDice: totalDice})
			}
		}

		//throw dice and save dice points
		for i := 0; i < len(players); i++ {
			players[i].Points = []int{}
			for j := 0; j < players[i].TotalDice; j++ {
				dicePoint := rand.Intn(6) + 1
				players[i].Points = append(players[i].Points, dicePoint)
			}
			fmt.Println(fmt.Sprintf("Player ke %v, Score %v , Point Dadu %v", players[i].Name, players[i].Score, players[i].Points))
		}

		//evaluate points and count dice
		for i := 0; i < len(players); i++ {
			//first process : remove dice with point 1,remove and add score dice with point 6, and keep the rest
			players[i].EvaluatedPoints = []int{}
			for j := 0; j < len(players[i].Points); j++ {
				if players[i].Points[j] == 1 {
					players[i].TotalDice--
				} else if players[i].Points[j] == 6 {
					players[i].Score++
					players[i].TotalDice--
				} else {
					players[i].EvaluatedPoints = append(players[i].EvaluatedPoints, players[i].Points[j])
				}
			}

			//second process : add dice with point one from beside player and still active
			for j := 0; j < CountPointOne(ActivePlayerPoints(i, players)); j++ {
				players[i].TotalDice++
				players[i].EvaluatedPoints = append(players[i].EvaluatedPoints, 1)
			}

			//third process : remove player who has no dice
			if players[i].TotalDice <= 0 {
				totalPlayer--
			}
			fmt.Println(fmt.Sprintf("Evaluasi Player ke %v, Score %v , Point Dadu %v", players[i].Name, players[i].Score, players[i].EvaluatedPoints))
		}

		//find the winner
		fmt.Println(fmt.Sprintf("Sisa Player Adalah %v", totalPlayer))
		if totalPlayer <= 1 {
			var highScorePlayer DiceGameModel
			for _, player := range players {
				if player.Score > highScorePlayer.Score && player.TotalDice == 0 {
					highScorePlayer = player
				}
			}
			fmt.Println(fmt.Sprintf("Pemenanganya Adalah Player %v Dengan Total Score %v", highScorePlayer.Name, highScorePlayer.Score))
			return
		}
		gameRound++
	}
}

func ActivePlayerPoints(order int, players []DiceGameModel) []int {
	//logic to find active player
	besidePlayer := order - 1
	for {
		if besidePlayer < 0 {
			besidePlayer = len(players) - 1
		}

		if players[besidePlayer].TotalDice != 0 {
			return players[besidePlayer].Points
		}

		if besidePlayer == order {
			return []int{}
		}
		besidePlayer--
	}
}

func CountPointOne(nums []int) int {
	var count int
	for _, num := range nums {
		if num == 1 {
			count++
		}
	}
	return count
}
