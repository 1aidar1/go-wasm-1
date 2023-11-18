package tictactoe

import (
	"errors"
	"fmt"
)

type Board struct {
	Over        bool
	Squares     [9]Square
	Turn        State
	MoveCounter uint
}
type Square struct {
	Id    int
	State State
}

func NewBoard() *Board {
	board := &Board{
		Squares: [9]Square{},
		Turn:    X,
	}
	for i := 0; i < len(board.Squares); i++ {
		board.Squares[i].Id = i
	}
	return board
}

func (b *Board) PlayInTerminal() State {
	for {
		// GAME END conditions
		winner, over := b.CalculateWinner()
		if over {
			return winner
		}

		// PLAY GAME
		var id int
		fmt.Scanf("%d", &id)
		if err := b.MakeMove(id); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s\n", b.ToString())
	}
}

func (b *Board) MakeMove(id int) error {
	if b.Over {
		return errors.New("game already over")
	}
	if id < 0 || id >= len(b.Squares) {
		return errors.New("id out of bounds of board")
	}
	if b.Squares[id].State != None {
		return errors.New("square not empty")
	}
	if b.Turn == None {
		return errors.New("who's turn is dat?")
	}
	// inc counter
	b.MoveCounter++
	// toggle turn
	b.Squares[id].State = b.Turn
	switch b.Turn {
	case X:
		b.Turn = O
	case O:
		b.Turn = X
	default:
		b.Turn = None
	}
	return nil
}

func (b *Board) CalculateWinner() (State, bool) {

	// Check rows and columns
	for i := 0; i < 3; i++ {
		if b.Squares[3*i].State == b.Squares[3*i+1].State &&
			b.Squares[3*i+1].State == b.Squares[3*i+2].State &&
			b.Squares[3*i].State != None {
			b.Over = true
			return b.Squares[3*i].State, true
		}

		if b.Squares[i].State == b.Squares[i+3].State &&
			b.Squares[i+3].State == b.Squares[i+6].State &&
			b.Squares[i].State != None {
			b.Over = true
			return b.Squares[i].State, true
		}
	}

	// Check diagonals
	if b.Squares[0].State == b.Squares[4].State &&
		b.Squares[4].State == b.Squares[8].State &&
		b.Squares[0].State != None {
		b.Over = true
		return b.Squares[0].State, true
	}

	if b.Squares[2].State == b.Squares[4].State &&
		b.Squares[4].State == b.Squares[6].State &&
		b.Squares[2].State != None {
		b.Over = true
		return b.Squares[2].State, true
	}
	// Check move counter
	if b.MoveCounter >= uint(len(b.Squares)) {
		b.Over = true
		return None, true
	}
	return None, false
}

func (b Board) ToString() string {
	var result string
	for i, square := range b.Squares {
		// Convert the square state to a readable symbol
		var symbol string
		switch square.State {
		case X:
			symbol = "X"
		case O:
			symbol = "O"
		default:
			symbol = " " // Represent Empty state with a space
		}

		// Append the symbol to the result
		result += symbol

		// Add line breaks after every third square
		if (i+1)%3 == 0 {
			result += "\n"
		} else {
			result += " | " // Add separators between squares
		}
	}
	return result
}
