package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./rps")

	if err != nil { panic(err) }

	scanner := bufio.NewScanner(file)

	if scanner.Err() != nil { panic(scanner.Err()) }

	rockPaperScissorStrategy := 0
	winLoseStrategy := 0

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		enemyMove := createFromABCString(line[0])
		possiblyMyMove := createFromXYZString(line[1])

		rockPaperScissorStrategyRound := possiblyMyMove.points

		if possiblyMyMove.losesToIndicator == enemyMove.indicator {
			rockPaperScissorStrategyRound += 0
		} else if possiblyMyMove.indicator == enemyMove.indicator {
			rockPaperScissorStrategyRound += 3
		} else {
			rockPaperScissorStrategyRound += 6
		}

		switch possiblyMyMove.possiblyAlternativeIndicator {
			case "X": winLoseStrategy += 0 + createFromABCString(enemyMove.winsFromIndicator).points
			case "Y": winLoseStrategy += 3 + enemyMove.points
			case "Z": winLoseStrategy += 6 + createFromABCString(enemyMove.losesToIndicator).points
		}

		rockPaperScissorStrategy += rockPaperScissorStrategyRound
	}

	fmt.Println(rockPaperScissorStrategy)
	fmt.Println(winLoseStrategy)

	err = file.Close()
	if err != nil { panic(err) }
}
