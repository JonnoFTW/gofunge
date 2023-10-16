package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"gofunge/gofunge"
)

func main() {
	showStep := flag.Bool("s", false, "Show board at each step")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: gofunge [-s] filename")
		return
	}
	filename := args[0]
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading %s: %v", filename, err)
	}

	i, err := gofunge.NewInterpreter(80, 25, string(content))
	if err != nil {
		log.Fatalf("Error making interpreter: %v", err)
	}
	input := bufio.NewScanner(os.Stdin)
	for {
		i.Step()
		if *showStep {
			exec.Command("clear")
			i.Show()
			input.Scan()
		}
	}
}
