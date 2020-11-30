package hangman

import (
	"fmt"
)

func DrawWelcome() {
	fmt.Println(`                       
 |_|  _. ._   _  ._ _   _. ._  
 | | (_| | | (_| | | | (_| | | 
              _|                                                                                        
	`)
}

// Draw the game
func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
		  +----+
		  |    |
		  O    |
		 /|\   |
		 / \   |
		       |
		=========
		`
	case 1:
		draw = `
		  +----+
		  |    |
		  O    |
		 /|\   |
		   \   |
		       |
		=========
		`
	case 2:
		draw = `
		  +----+
		  |    |
		  O    |
		 /|\   |
		       |
		       |
		=========
		`
	case 3:
		draw = `
		  +----+
		  |    |
		  O    |
		  |\   |
		       |
		       |
		=========
		`
	case 4:
		draw = `
		  +----+
		  |    |
		  O    |
		  |    |
		       |
		       |
		=========
		`
	case 5:
		draw = `
		  +----+
		  |    |
		  O    |
		       |
		       |
		       |
		=========
		`
	case 6:
		draw = `
		  +----+
		  |    |
		       |
		       |
		       |
		       |
		=========
		`
	case 7:
		draw = `






		=========
		`
	case 8:
		draw = `



		


	
		`
	}
	fmt.Println(draw)
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)
	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {    
	case "goodGuess":
		fmt.Println("Good Guess !")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' was already used\n", guess)
	case "badGuess":
		fmt.Printf("Bad Guess, %s is not in the word\n", guess)
	case "lost":
		fmt.Print("You LOST ! The word was: ")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("You WON ! The word was: ")
		drawLetters(g.Letters)
	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v", c)
	}
	fmt.Println()
}
