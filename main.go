package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
	Yellow = "\033[33m"
	Blue  = "\033[34m"
	Black = "\033[30m" // <- preto (vai "sumir" no fundo preto)
)

const size = 9

type Board [size][size]int

func (b *Board) Print() {
	for i := 0; i < size; i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("------+-------+------")
		}
		for j := 0; j < size; j++ {
			if j%3 == 0 && j != 0 {
				fmt.Print("| ")
			}

			val := b[i][j]
			if val == 0 {
				// pinta o 0 de preto pra "sumir" no fundo
				fmt.Printf("%s0 %s", Black, Reset)
			} else {
				fmt.Printf("%d ", val)
			}
		}
		fmt.Println()
	}
}

func isSafe(b *Board, row, col, num int) bool {
	// Linha
	for x := 0; x < size; x++ {
		if b[row][x] == num {
			return false
		}
	}

	// Coluna
	for y := 0; y < size; y++ {
		if b[y][col] == num {
			return false
		}
	}

	// Bloco 3x3
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if b[startRow+r][startCol+c] == num {
				return false
			}
		}
	}

	return true
}

func (b *Board) findEmpty() (int, int, bool) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if b[i][j] == 0 {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func randPerm1to9() []int {
	nums := rand.Perm(9) // 0..8
	for i := range nums {
		nums[i]++ // vira 1..9
	}
	return nums
}

// Backtracking preenchendo com números aleatórios
func (b *Board) solveRandom() bool {
	row, col, found := b.findEmpty()
	if !found {
		return true // sem espaços vazios: resolvido
	}

	nums := randPerm1to9()
	for _, num := range nums {
		if isSafe(b, row, col, num) {
			b[row][col] = num

			if b.solveRandom() {
				return true
			}

			// backtrack
			b[row][col] = 0
		}
	}

	return false
}

// Remove 'holes' números aleatoriamente (sem se preocupar com solução única)
func (b *Board) removeCells(holes int) {
	attempts := holes
	for attempts > 0 {
		row := rand.Intn(size)
		col := rand.Intn(size)

		if b[row][col] != 0 {
			b[row][col] = 0
			attempts--
		}
	}
}

func GenerateSudoku(holes int) {
	rand.Seed(time.Now().UnixNano())

	var board Board

	// Gera um tabuleiro completo válido
	if !board.solveRandom() {
		fmt.Println(Red, "Falha ao gerar Sudoku :(", Reset)
		return
	}

	fmt.Println(Green, "Tabuleiro completo:", Reset)
	board.Print()

	// Copia o board completo para gerar o puzzle
	puzzle := board
	puzzle.removeCells(holes) // quantidade de casas vazias (ajusta a 'dificuldade')

	fmt.Printf(Yellow+"\nSudoku com espaços vazios):\n"+Reset)
	puzzle.Print()
}

func main() {
	var resposta string

	fmt.Print("Você quer gerar um Sudoku? (s/n): ")
	fmt.Scanln(&resposta)

	if resposta == "s" || resposta == "S" {
		fmt.Println(Blue, "Beleza, vou gerar o Sudoku!", Reset)
		GenerateSudoku(40) // aqui você controla a dificuldade
	} else {
		fmt.Println("Tranquilo, não vou gerar nada.")
	}
}
