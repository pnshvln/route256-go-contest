package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var usersCount int
	var pairsCount int
	fmt.Fscan(in, &usersCount, &pairsCount)

	users := make(map[int][]int)
	for i := 0; i < pairsCount; i++ {
		var user1 int
		var user2 int
		fmt.Fscan(in, &user1, &user2)
		users[user1] = append(users[user1], user2)
		users[user2] = append(users[user2], user1)
	}

	for i := 1; i <= usersCount; i++ {
		friends, exist := users[i]
		if !exist {
			fmt.Fprintln(out, 0)
			continue
		}

		freq := make(map[int]int)

		for _, friend := range friends {
			fof := users[friend]
			for _, ff := range fof {
				if ff != i && !inSlice(ff, friends) {
					freq[ff] += 1
				}
			}
		}

		max := -1
		var recommend []int = make([]int, 0)
		for k, v := range freq {
			if v == max {
				recommend = append(recommend, k)
			}

			if v > max {
				max = v
				recommend = make([]int, 0)
				recommend = append(recommend, k)
			}
		}

		if len(recommend) <= 0 {
			fmt.Fprintln(out, 0)
		} else {
			sort.Slice(recommend, func(j, k int) bool {
				return recommend[j] < recommend[k]
			})
			for i := 0; i < len(recommend); i++ {
				fmt.Fprintf(out, "%d ", recommend[i])
			}
			fmt.Fprintln(out)
		}
	}
}

func inSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
