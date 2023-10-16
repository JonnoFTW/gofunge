package gofunge

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	color "github.com/fatih/color"
)

type Direction int

const (
	Up    Direction = iota
	Down  Direction = iota
	Left  Direction = iota
	Right Direction = iota
)

type Interpeter struct {
	stack        Stack
	px           int // program counter x pos
	py           int // program counter y pos
	board        [][]byte
	stringMode   bool
	skippingMode bool
	direction    Direction
}

func NewInterpreter(width int, height int, source string) (*Interpeter, error) {
	board := make([][]byte, height)
	for i := range board {
		board[i] = make([]byte, width)
	}
	lines := strings.Split(source, "\n")
	for y := 0; y < len(lines); y++ {
		line := lines[y]
		for x := 0; x < len(line); x++ {
			board[y][x] = line[x]
		}
	}

	return &Interpeter{Stack{}, 0, 0, board, false, false, Right}, nil
}

func (i *Interpeter) Step() {
	// execute the character
	c := i.board[i.py][i.px]
	if i.stringMode {
		if c == '"' {
			i.stringMode = false
		} else {
			i.stack.Push(int(c))
		}
	} else {
		if i.skippingMode {
			i.skippingMode = false
		} else {
			switch c {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				i.stack.Push(int(c) - 48)
			case '+':
				i.stack.Push(i.stack.Pop() + i.stack.Pop())
			case '*':
				i.stack.Push(i.stack.Pop() * i.stack.Pop())
			case '-':
				a := i.stack.Pop()
				b := i.stack.Pop()
				i.stack.Push(b - a)
			case '/':
				a := i.stack.Pop()
				b := i.stack.Pop()
				i.stack.Push(b / a)
			case '%':
				a := i.stack.Pop()
				b := i.stack.Pop()
				i.stack.Push(b % a)
			case '!':
				a := i.stack.Pop()
				if a == 0 {
					i.stack.Push(1)
				} else {
					i.stack.Push(0)
				}
			case '`':
				a := i.stack.Pop()
				b := i.stack.Pop()
				if b > a {
					i.stack.Push(1)
				} else {
					i.stack.Push(0)
				}
			case '>':
				i.direction = Right
			case '<':
				i.direction = Left
			case '^':
				i.direction = Up
			case 'v':
				i.direction = Down
			case '?':
				i.direction = Direction(rand.Int() % 4)
			case '_':
				a := i.stack.Pop()
				if a == 0 {
					i.direction = Right
				} else {
					i.direction = Left
				}
			case '|':
				a := i.stack.Pop()
				if a == 0 {
					i.direction = Down
				} else {
					i.direction = Up
				}
			case '"':
				i.stringMode = !i.stringMode
			case ':':
				i.stack.Push(i.stack.Peek())
			case '\\':
				a := i.stack.Pop()
				b := i.stack.Pop()
				i.stack.Push(a)
				i.stack.Push(b)
			case '$':
				i.stack.Pop()
			case '.':
				fmt.Printf("%d ", i.stack.Pop()) // print stack item as int with space
			case ',':
				fmt.Print(string(i.stack.Pop())) // print stack item as ascii
			case '#':
				// bridge over
				i.skippingMode = true
			case 'p':
				// pop y x v . Set character at x,y to ascii value of v
				y := i.stack.Pop()
				x := i.stack.Pop()
				v := i.stack.Pop()
				i.board[y][x] = byte(v)
			case 'g':
				// pop y and x then push value at those coords to stack
				y := i.stack.Pop()
				x := i.stack.Pop()
				i.stack.Push(int(i.board[y][x]))
			case '&':
				// ask for user number input and push
				fmt.Print("># ")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
			case '~':
				fmt.Print((">A "))
				// ask for char and push ascii value
			case '@':
				// end program
				os.Exit(0)
			default:
			}
		}
	}
	// adjust program counters accordingly
	switch i.direction {
	case Up:
		i.py--
	case Down:
		i.py++
	case Left:
		i.px--
	case Right:
		i.px++
	}

}
func (i *Interpeter) Show() {
	for y := 0; y < len(i.board); y++ {
		row := i.board[y]
		for x := 0; x < len(row); x++ {
			if x == i.px && y == i.py {
				// print the pc in green background
				c := color.New(color.BgGreen)
				c.Print(string((i.board[y][x])))
			} else {
				// print everything else normally
				fmt.Print(string(i.board[y][x]))
			}
		}
		fmt.Println()
	}
}
