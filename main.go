package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Fprintln(os.Stderr, "Error loafing env file: "+err.Error())
		os.Exit(1)
	}
}

func main() {
	loadEnv()
	// initChatGPT()
	p := tea.NewProgram(initModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
