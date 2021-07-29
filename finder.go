package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/domino14/macondo/gaddag"
)

var answerSet map[string]int

func printAnswers() {

	type Answer struct {
		score int
		word  string
	}
	printableAnswers := map[int][]Answer{}

	for w, s := range answerSet {
		l := len(w)
		if _, ok := printableAnswers[l]; !ok {
			printableAnswers[l] = []Answer{}
		}
		printableAnswers[l] = append(printableAnswers[l], Answer{score: s, word: w})
	}

	for i := 15; i >= 2; i-- {
		if answers, ok := printableAnswers[i]; ok {
			fmt.Printf("%d letters (%d answers)\n", i, len(answers))
			fmt.Println("-------------------------")
			sort.Slice(answers, func(i, j int) bool {
				return answers[i].word < answers[j].word
			})
			for _, a := range answers {
				fmt.Printf("%s (%d) ", a.word, a.score)
			}
			fmt.Println()
			fmt.Println()
		}
	}
	fmt.Printf("%d total answers\n", len(answerSet))
}

// wuggler finds all the words in a square board passed in. The length
// of the board string must be a perfect square (it is converted if not,
// depending on the round)
func wuggler(dawg *gaddag.SimpleDawg, board []rune, round int) {
	printBoard(board, round)
	boardDim := roundToDim(round)
	answerSet = make(map[string]int)
	for idx := range board {
		if board[idx] == ' ' {
			continue
		}
		newBoard := removeLetter(idx, board)
		multiplier := boardBonus(idx, round)
		findWords(dawg, idx, newBoard, []rune{board[idx]},
			boardDim, multiplier, round)
	}
}

// Convert the string into a board shape depending on the round
func convertBoard(board string, round int) ([]rune, error) {
	switch round {
	case 1:
		if len(board) != 16 {
			return nil, errors.New("need 16 letters for round 1")
		}
		return []rune(board), nil
	case 2:
		if len(board) != 24 {
			return nil, errors.New("need 24 letters for round 2")
		}
		strIdx := 0
		newStr := ""
		for y := 0; y < 6; y++ {
			for x := 0; x < 6; x++ {
				if x == 0 || x == 5 {
					if y < 2 || y > 3 {
						newStr += " "
						continue
					}
				}
				if x == 1 || x == 4 {
					if y < 1 || y > 4 {
						newStr += " "
						continue
					}
				}
				newStr += string(board[strIdx])
				strIdx++
			}
		}
		return []rune(newStr), nil
	case 3:
		if len(board) != 28 {
			return nil, errors.New("need 28 letters for round 3")
		}
		strIdx := 0
		newStr := ""
		for y := 0; y < 6; y++ {
			for x := 0; x < 6; x++ {
				if x == 0 || x == 1 {
					if y > 3 {
						newStr += " "
						continue
					}
				}
				if x == 4 || x == 5 {
					if y < 2 {
						newStr += " "
						continue
					}
				}
				newStr += string(board[strIdx])
				strIdx++
			}
		}
		return []rune(newStr), nil

	case 4:
		if len(board) != 32 {
			return nil, errors.New("need 32 letters for round 4")
		}
		strIdx := 0
		newStr := ""
		for y := 0; y < 6; y++ {
			for x := 0; x < 6; x++ {
				if x == 2 || x == 3 {
					if y == 2 || y == 3 {
						newStr += " "
						continue
					}
				}
				newStr += string(board[strIdx])
				strIdx++
			}
		}
		return []rune(newStr), nil
	}
	return nil, errors.New("round not found")
}

func roundToDim(round int) int {
	var dim int
	if round == 1 {
		dim = 4
	} else {
		dim = 6
	}
	return dim
}

func printBoard(board []rune, round int) {
	var Reset = "\033[0m"
	var Red = "\u001b[31m"
	var Blue = "\u001b[34m"
	dim := roundToDim(round)
	fmt.Println("-------------")
	for y := 0; y < dim; y++ {
		for x := 0; x < dim; x++ {
			color := ""
			reset := ""
			idx := xyToIndex(x, y, dim)
			bonus := boardBonus(idx, round)
			if bonus == 2 {
				color = Blue
				reset = Reset
			} else if bonus == 3 {
				color = Red
				reset = Reset
			}
			letter := string(board[idx])
			afterSpace := " "
			if letter == "Q" {
				letter = "Qu"
				afterSpace = ""
			}
			fmt.Printf("%v%v%v%v", color, letter, reset, afterSpace)
		}
		fmt.Printf("\n")
	}
	fmt.Println("-------------")
}

func stringToFind(possibleWord []rune) string {
	// Replace Q with QU everywhere.
	newStr := strings.Replace(string(possibleWord), "Q", "QU", -1)
	return newStr
}

func findWords(dawg *gaddag.SimpleDawg, idx int, board []rune,
	possibleWord []rune, boardDim int, multiplier int, round int) {
	// possibleWord is actually a prefix, perfect for dawg. Check
	// if this prefix is in the dawg. If it's not, it's time to prune
	// this branch.
	if !gaddag.FindPrefix(dawg, stringToFind(possibleWord)) {
		return
	}
	allowable := allowableIndices(idx, boardDim)
	for _, newIdx := range allowable {
		if board[newIdx] == ' ' {
			continue
		}
		newBoard := removeLetter(newIdx, board)
		newMultiplier := boardBonus(newIdx, round)
		findWords(dawg, newIdx, newBoard, append(possibleWord, board[newIdx]),
			boardDim, multiplier*newMultiplier, round)
	}
	if gaddag.FindWord(dawg, stringToFind(possibleWord)) {
		addPlay([]rune(stringToFind(possibleWord)), multiplier)
	}
}

// removeLetter removes the letter at idx from board.
func removeLetter(idx int, board []rune) []rune {
	var newBoard = make([]rune, len(board))
	for i, letter := range board {
		if i != idx {
			newBoard[i] = letter
		} else {
			newBoard[i] = ' '
		}
	}
	return []rune(newBoard)
}

// allowableIndices finds what indices in the board are allowed to be
// reached from idx
func allowableIndices(idx int, boardDim int) []int {
	y := idx / boardDim
	x := idx % boardDim
	var indices []int
	// There are eight directions
	// up down left right
	if x+1 < boardDim {
		indices = append(indices, xyToIndex(x+1, y, boardDim))
	}
	if x > 0 {
		indices = append(indices, xyToIndex(x-1, y, boardDim))
	}
	if y+1 < boardDim {
		indices = append(indices, xyToIndex(x, y+1, boardDim))
	}
	if y > 0 {
		indices = append(indices, xyToIndex(x, y-1, boardDim))
	}
	// diagonals
	if x+1 < boardDim && y+1 < boardDim {
		indices = append(indices, xyToIndex(x+1, y+1, boardDim))
	}
	if x > 0 && y > 0 {
		indices = append(indices, xyToIndex(x-1, y-1, boardDim))
	}
	if x > 0 && y+1 < boardDim {
		indices = append(indices, xyToIndex(x-1, y+1, boardDim))
	}
	if x+1 < boardDim && y > 0 {
		indices = append(indices, xyToIndex(x+1, y-1, boardDim))
	}
	return indices
}

func xyToIndex(x int, y int, boardDim int) int {
	return y*boardDim + x
}

func indexToXY(idx int, boardDim int) (int, int) {
	return idx % boardDim, idx / boardDim
}

func addPlay(word []rune, multiplier int) {
	newScore := score(word, multiplier)

	if oldScore, ok := answerSet[string(word)]; !ok || newScore > oldScore {
		answerSet[string(word)] = newScore
	}

}
