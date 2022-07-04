package main

import (
	"bufio"
	"fmt"
	"math"
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

		progs := make([]int, count)

		for j := 0; j < count; j++ {
			fmt.Fscan(in, &progs[j])
		}

		for k := 0; k < count; k++ {
			if progs[k] < 0 {
				continue
			}

			index1 := k + 1
			index2 := k + 1

			first := progs[k]
			diff := math.MaxInt

			for j := 1; j < count; j++ {
				if progs[j] < 0 || j == k {
					continue
				}
				d := Abs(first - progs[j])
				if d < diff {
					diff = d
					index2 = j + 1
				}
			}

			progs[index1-1] = -1
			progs[index2-1] = -1
			fmt.Fprintln(out, index1, index2)
		}
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
