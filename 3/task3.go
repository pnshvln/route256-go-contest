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

	var usersCount int
	var commandsCount int
	fmt.Fscan(in, &usersCount, &commandsCount)

	msgId := 1
	globalMsgId := 0
	users := make(map[int]int)

	for i := 0; i < commandsCount; i++ {
		var commandType int
		var userId int
		fmt.Fscan(in, &commandType, &userId)

		if commandType == 1 {
			if userId == 0 {
				globalMsgId = msgId
			} else {
				users[userId] = msgId
			}
			msgId += 1
		} else {
			value, exist := users[userId]
			if exist {
				if globalMsgId > value {
					fmt.Fprintln(out, globalMsgId)
				} else {
					fmt.Fprintln(out, value)
				}
			} else {
				if globalMsgId > 0 {
					fmt.Fprintln(out, globalMsgId)
				} else {
					fmt.Fprintln(out, 0)
				}
			}
		}
	}
}
