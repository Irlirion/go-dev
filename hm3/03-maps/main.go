package main

import "fmt"

func StudentScores() map[string]int {
	scores := make(map[string]int)

	scores["Анна"] = 3
	scores["Борис"] = 0
	scores["Вера"] = 5

	scores["Анна"] += 2
	delete(scores, "Вера")

	return scores
}

func main() {
	scores := StudentScores()

	fmt.Println(scores["Анна"]) // 5

	score, ok := scores["Борис"]
	fmt.Println(score, ok) // 0 true

	score, ok = scores["Глеб"]
	fmt.Println(score, ok) // 0 false

	fmt.Println(len(scores)) // 2
}
