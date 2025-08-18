package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	clnPath, err := os.UserHomeDir()
	clnPath = clnPath + "/cln"
	if err != nil {
		panic(err)
	}

	if !pathExists(clnPath) {
		os.Mkdir(clnPath, 0755)
	}

	args := os.Args

	noteText := args[1]

	today := time.Now().Local().Format("02-01-2006")

	notesPath := clnPath + "/" + today + ".md"

	file, err := os.OpenFile(notesPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	now := time.Now().Local().Format("15:04:05")

	note := now + " - " + noteText + "\n"

	fmt.Println(note)
	file.WriteString(note)
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
