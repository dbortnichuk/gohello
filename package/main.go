package main

import (
	"github.com/dbortnichuk/gohello/package/world"
	"fmt"
)


func main() {

	fmt.Println("export from different package", world.GetStartingLevel(), world.StartingRoom)
	fmt.Println("export from different package unicode name", world.Змінна)
	fmt.Println("export from different package mixed name", world.ОтриматиЗміннаWWW())

	fmt.Println("export from different package which uses another import", world.CurrentStore())
	fmt.Println("export from same package but different file", maxLevel)
	fmt.Println("export function from different package", world.IsStartingLevel(2, "unused param"))


}
