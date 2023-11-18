package main

import (
	"fmt"
	"syscall/js"

	"github.com/1aidar/wasm-app/tictactoe"
)

func MakeMoveJs(board *tictactoe.Board) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Check if there are enough arguments (for the square ID)
		if len(args) == 0 {
			return map[string]interface{}{
				"error": "No square ID provided",
			}
		}

		// Get the square ID from the first argument
		id := args[0].Int()
		// Make the move
		if err := board.MakeMove(id); err != nil {
			return map[string]interface{}{
				"error": err.Error(),
			}
		}

		fmt.Printf("%+v\n", board.ToString())

		// Check for a winner or if the game is over
		state, over := board.CalculateWinner()
		// Return the current state of the game
		return map[string]interface{}{
			"board":    board.ToString(),
			"winner":   state.ToString(),
			"gameOver": over,
			"error":    nil,
		}
	})
}

func main() {

	ch := make(chan struct{}, 0)
	fmt.Printf("is this console log?\n")
	game := tictactoe.NewBoard()
	js.Global().Set("MakeMoveJs", MakeMoveJs(game))
	// js.Global().Set("Restart", RestartJs(game))

	<-ch
}
