package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	file, err := os.OpenFile("readme.md", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error When Opening File: ", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString("auto commit by khairulharu sheeexzzz")

	if err != nil {
		fmt.Println("Erorr writing to file: ", err)
	}

	cmdAdd := exec.Command("git", "add", ".")

	cmdAdd.Dir = currentDirectory

	outputAdd, err := cmdAdd.Output()
	if err != nil {
		fmt.Println("Error running command:", err)
		return
	}

	fmt.Println(string(outputAdd))

	cmdCommit := exec.Command("git", "commit", "-m", "auto commit bro made by khairulharu sigma")

	cmdCommit.Dir = currentDirectory

	outputCommit, err := cmdCommit.Output()
	if err != nil {
		fmt.Println("Error running command:", err)
		return
	}

	fmt.Println(string(outputCommit))

	cmdPush := exec.Command("git", "push", "origin", "master")

	cmdPush.Dir = currentDirectory

	outputPush, err := cmdPush.Output()
	if err != nil {
		fmt.Println("Error running command:", err)
		return
	}

	fmt.Println(string(outputPush))

	fmt.Println("Auto Commit To")
}
