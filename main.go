package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gamingdiamond982/randomfish/board"
)

var stdin chan string = make(chan string, 100)
var stdout chan string = make(chan string, 100)

const startposfen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
const ready = true

func handleStdio() {
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			msg, _ := reader.ReadString('\n')
			stdin <- strings.Replace(msg, "\n", "", -1)
		}
	}()

	for msg := range stdout {
		fmt.Println(msg)
	}

}

func main() {
	go handleStdio()
	var brd board.Board

	for msg := range stdin {
		args := strings.Split(msg, " ")
		switch args[0] {
		case "uci":
			stdout <- "id randomfish"
			stdout <- "uciok"
		case "isready":
			go func() {
				for !ready {
					time.Sleep(time.Second * 1)
				}
				stdout <- "readyok"
			}()

		case "quit":
			return

		case "ucinewgame":
			//TODO: do something here idk

		case "position":
			fen := strings.Replace(msg, args[0], "", -1)
			if fen == " startpos" {
				fen = startposfen
			}
			brd = board.CreateBoard(fen)
		case "show":
			stdout <- fmt.Sprint(brd)

		default:
			stdout <- fmt.Sprintf("Unknown command: %v", args[0])

		}
	}

}
