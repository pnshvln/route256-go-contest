package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var count int
		fmt.Fscan(in, &count)

		finished := make(map[int]int)
		lastTask := -1

		isCorrect := true

		for k := 0; k < count; k++ {
			var task int
			fmt.Fscan(in, &task)

			if !isCorrect {
				continue
			}

			if lastTask == -1 {
				lastTask = task
				continue
			}

			if lastTask != task {
				_, exist := finished[task]

				if exist {
					isCorrect = false
				} else {
					finished[lastTask] = -1
				}
				lastTask = task
			}
		}

		if isCorrect {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}

	}
}
