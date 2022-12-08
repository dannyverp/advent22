package main

type moveCreator interface {
	create(indicator string) move
}

type move struct {
	points int
	indicator string
	losesToIndicator string
	winsFromIndicator string
	possiblyAlternativeIndicator string
	name string
}

func createFromABCString(indicator string) move {
	if indicator == "A"{
		return move{1, "A", "B","C","X", "ROCK"}
	} else if indicator == "B" {
		return move{2,"B", "C","A","Y", "PAPER"}
	} else if indicator == "C" {
		return move{3,"C", "A","B","Z", "SCISSOR"}
	}
	panic("Hold on, that type doesn't exist")
}

func createFromXYZString(indicator string) move {
	var result move
	switch indicator {
		case "X": result = createFromABCString("A")
		case "Y": result = createFromABCString("B")
		case "Z": result = createFromABCString("C")
	}
	return result
}