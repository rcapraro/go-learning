package hangman

import (
	"fmt"
	"strings"
)

// Game represents a hangman game
type Game struct {
	State        string   //Game State
	Letters      []string //Letters in the word to find
	FoundLetters []string //Good guesses
	UsedLetters  []string //Used                                               letters
	TurnsLeft    int      //Remaining attempts
}

//New creates a new hangman game
func New(turns int, word string) (*Game, error) {
	if len(word) < 1 {
		return nil, fmt.Errorf("Word '%s' must be at least one character. got %v", word, len(word))
	}
	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}
	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}
	return g, nil
}

// MakeAGuess makes a guess
func (g *Game) MakeAGuess(guess string) {
	guess = strings.ToUpper(guess)
	switch g.State {
	case "won", "lost":
		return
	}
	if letterInWord(guess, g.UsedLetters) {
		g.State = "alreadyGuessed"
	} else if letterInWord(guess, g.Letters) {
		g.State = "goodGuess"
		g.revealLetter(guess)
		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}
	} else {
		g.looseTurn(guess)
		if hasLost(g) {
			g.State = "lost"
		}
	}
}

func letterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}

func (g *Game) looseTurn(guess string) {
	g.State = "badGuess"
	g.UsedLetters = append(g.UsedLetters, guess)
	g.TurnsLeft--
}

func hasWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}
	return true
}

func hasLost(g *Game) bool {
	if g.TurnsLeft <= 0 {
		return true
	}
	return false
}

func (g *Game) revealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess
		}
	}
}
