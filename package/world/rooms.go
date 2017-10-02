package world

import "github.com/dbortnichuk/gohello/package/world1"

const StartingRoom = "lobby"
const startingLevel  = 1

var Змінна = "variable"

func GetStartingLevel() int {
	return startingLevel
}

func IsStartingLevel(level int, param string ) bool {
	return level == startingLevel
}

func getVar() string {
	return Змінна
}

func ОтриматиЗміннаWWW() string {
	return Змінна
}

func CurrentStore() int {
	return world.Square
	}