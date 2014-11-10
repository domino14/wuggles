package main

import (
	"fmt"
	"github.com/domino14/macondo/gaddag"
)

var answerSet map[string]bool

// wuggler finds all the words in a square board passed in. The length
// of the board string must be a perfect square.
func wuggler(g gaddag.SimpleGaddag, board string) {
	fmt.Println(board)
	if len(board) != 16 {
		fmt.Println("Only 16-letter boards are supported.")
		return
	}
	boardDim := uint8(4)
	answerSet = make(map[string]bool)
	for idx, _ := range board {
		newBoard := removeLetter(uint8(idx), board)
		findWords(g, uint8(idx), newBoard, string(board[idx]), boardDim)
	}
	fmt.Println(answerSet)
	fmt.Println(len(answerSet), "answers")
}

func findWords(g gaddag.SimpleGaddag, idx uint8, board string,
	possibleWord string, boardDim uint8) {
	// possibleWord is actually a prefix, perfect for GADDAG. Check
	// if this prefix is in the GADDAG. If it's not, it's time to prune
	// this branch.
	if !gaddag.FindPrefix(g, possibleWord) {
		return
	}
	allowable := allowableIndices(idx, boardDim)
	for _, newIdx := range allowable {
		if board[newIdx] == ' ' {
			continue
		}
		newBoard := removeLetter(uint8(newIdx), board)
		findWords(g, newIdx, newBoard, possibleWord+string(board[newIdx]),
			boardDim)
	}
	if gaddag.FindWord(g, possibleWord) {
		addPlay(possibleWord)
	}
}

// removeLetter removes the letter at idx from board.
func removeLetter(idx uint8, board string) string {
	// var boardCopy string
	// if idx == 0 {
	// 	boardCopy = string(' ') + board[idx+1:]
	// } else {
	// 	boardCopy = board[0:idx] + string(' ') + board[idx+1:]
	// }
	// return boardCopy
	var newBoard string
	newBoard = ""
	for i, letter := range board {
		if uint8(i) != idx {
			newBoard += string(letter)
		} else {
			newBoard += string(' ')
		}
	}
	return newBoard
}

// allowableIndices finds what indices in the board are allowed to be
// reached from idx
func allowableIndices(idx uint8, boardDim uint8) []uint8 {
	y := idx / boardDim
	x := idx % boardDim
	var indices []uint8
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

func xyToIndex(x uint8, y uint8, boardDim uint8) uint8 {
	return y*boardDim + x
}

func addPlay(word string) {
	answerSet[word] = true
}
