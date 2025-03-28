package core

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
	"github.com/rushikeshg25/cool-wire/wire"
)

var (
	boldGreen = color.New(color.FgGreen, color.Bold).SprintFunc()
	boldRed   = color.New(color.FgRed, color.Bold).SprintFunc()
	boldBlue  = color.New(color.FgBlue, color.Bold).SprintFunc()
)

func Run(host string, port int) {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:      fmt.Sprintf("cool@%s:%s> ", boldBlue(host), boldBlue(port)),
		HistoryFile: os.ExpandEnv("$HOME/.dicedb_history"),
	})
	if err != nil {
		fmt.Printf("%s failed to initialize readline: %v\n", boldRed("ERR"), err)
		return
	}
	defer rl.Close()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		<-sigChan
		fmt.Println("\nreceived interrupt. exiting...")
		os.Exit(0)
	}()

	for {
		input, err := rl.Readline()
		if err != nil {
			break
		}
		if input == "exit" {
			return
		}
		if input == "" {
			continue
		}

		query := parseQuery(input)
		if len(query) == 0 {
			continue
		}

		q := wire.Query{
			Query: query,
		}
	}
}

func parseQuery(input string) string {
	//TODO: Add more validations
	return strings.TrimSpace(input)
}
