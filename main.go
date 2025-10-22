package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	// Set range tanggal
	startDate := time.Date(2025, 10, 22, 10, 0, 0, 0, time.Local)
	endDate := time.Date(2025, 10, 24, 18, 0, 0, 0, time.Local)

	// Random commits per hari (misal 1-5 commits per hari)
	minCommitsPerDay := 1
	maxCommitsPerDay := 5

	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// Loop untuk setiap hari di range
	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		// Random berapa kali commit di hari ini
		commitsToday := rand.Intn(maxCommitsPerDay-minCommitsPerDay+1) + minCommitsPerDay

		fmt.Printf("Tanggal %s: %d commits\n", d.Format("2006-01-02"), commitsToday)

		// Buat commits sebanyak yang di-random
		for i := 0; i < commitsToday; i++ {
			// Random jam dalam sehari (misal antara jam 9 pagi - 6 sore)
			randomHour := rand.Intn(9) + 9 // 9-17
			randomMinute := rand.Intn(60)
			randomSecond := rand.Intn(60)

			commitDate := time.Date(d.Year(), d.Month(), d.Day(),
				randomHour, randomMinute, randomSecond, 0, time.Local)

			// Tulis ke file
			file, err := os.OpenFile("readme.md", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				fmt.Println("Error When Opening File:", err)
				continue
			}

			commitMsg := fmt.Sprintf("auto commit %s #%d\n",
				commitDate.Format("2006-01-02 15:04:05"), i+1)
			_, err = file.WriteString(commitMsg)
			file.Close()

			if err != nil {
				fmt.Println("Error writing to file:", err)
				continue
			}

			// Git add
			cmdAdd := exec.Command("git", "add", ".")
			cmdAdd.Dir = currentDirectory
			if err := cmdAdd.Run(); err != nil {
				fmt.Println("Error git add:", err)
				continue
			}

			// Git commit dengan custom date
			dateStr := commitDate.Format("Mon Jan 2 15:04:05 2006 -0700")
			cmdCommit := exec.Command("git", "commit",
				"--date", dateStr,
				"-m", fmt.Sprintf("auto commit %s", commitDate.Format("2006-01-02 15:04:05")))
			cmdCommit.Dir = currentDirectory

			// Set environment variable untuk committer date
			cmdCommit.Env = append(os.Environ(),
				fmt.Sprintf("GIT_COMMITTER_DATE=%s", dateStr))

			output, err := cmdCommit.CombinedOutput()
			if err != nil {
				fmt.Println("Error git commit:", err)
				fmt.Println(string(output))
				continue
			}

			fmt.Printf("✓ Committed: %s\n", commitDate.Format("2006-01-02 15:04:05"))
		}
	}

	// Push semua commits sekaligus
	fmt.Println("\nPushing to remote...")
	cmdPush := exec.Command("git", "push", "origin", "master")
	cmdPush.Dir = currentDirectory
	outputPush, err := cmdPush.CombinedOutput()
	if err != nil {
		fmt.Println("Error git push:", err)
		fmt.Println(string(outputPush))
		return
	}

	fmt.Println("Auto Commit Successfully Run On Your Project")
}
