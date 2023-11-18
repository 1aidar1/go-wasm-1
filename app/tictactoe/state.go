package tictactoe

//State

type State int

func (s State) ToString() string {
	switch s {
	case 0:
		return "None"
	case 1:
		return "X"
	case 2:
		return "O"
	default:
		return ""
	}
}

const (
	None State = iota
	X
	O
)
