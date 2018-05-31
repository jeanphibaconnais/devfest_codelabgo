package main

import (
	"fmt"
	"os"
	"errors"
)

func main() {
	var message string    	// manual type déclaration

	// var args []string
	args := os.Args[1:] 		// type inference
	nbParam := len(args)

	fmt.Sprintf("\nExercice 2 : %d", nbParam)

	if nbParam == 0 {
		message = "décollage !"
	} else if nbParam == 1 {
		message = args[1]
	} else {
		err := errors.New("Paramètres invalides")
		fmt.Print(err)
		//panic(fmt.Sprintf("Mayday !"))
	}


	fmt.Print(message + " ---- \n\n")

}
