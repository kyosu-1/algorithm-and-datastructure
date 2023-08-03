package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// N の読み取り
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	// 配列 A の読み取り
	A := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.Atoi(scanner.Text())
	}

	// 累積和
	S := make([]int, N+1)
	for i := 0; i < N; i++ {
		S[i+1] = S[i] + A[i]
	}

	var Q int
	scanner.Scan()
	Q, _ = strconv.Atoi(scanner.Text()) // この行を追加

	for i := 0; i < Q; i++ {
		var L, R int
		scanner.Scan()
		L, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		R, _ = strconv.Atoi(scanner.Text())
		diff := (R - L + 1) - 2*(S[R] - S[L-1])
		if diff > 0 {
			fmt.Println("lose")
		} else if diff == 0 {
			fmt.Println("draw")
		} else {
			fmt.Println("win")
		}
	}
}
