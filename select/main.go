package main

import (
	"fmt"
	"math/rand"
)

// 集合 A の i 番目の順序統計量を求める
func randomizedSelect(A []int, p int, r int, i int) int {
	if p == r {
		fmt.Println("A. 要素数が 1")
		return A[p] // 要素数が 1
	}

	q := randomizedPartion(A, p, r)
	k := q - p + 1

	printArrayWithPivot(A, p, q, r)
	// fmt.Println("A:", A, "p:", p, "q:", q, "r:", r, "i:", i, "k:", k)
	if i == k {
		fmt.Println(" --- A. ピボットが答え")
		return A[q] // ピボットが i 番目の順序統計量
	} else if i < k {
		fmt.Println(" --- 再帰a. ピボット以下の部分配列のみを探索")
		return randomizedSelect(A, p, q-1, i) // ピボット以下の部分配列のみを探索
	}
	fmt.Println(" --- 再帰b. ピボット以上の部分配列のみを探索")
	return randomizedSelect(A, q+1, r, i-k) // ピボット以上の部分配列のみを探索
}

func randomizedPartion(A []int, p int, r int) int {
	i := rand.Intn(r-p) + p
	A[r], A[i] = A[i], A[r]
	return partion(A, p, r)
}

func partion(A []int, p int, r int) int {
	x := A[r]
	i := p - 1

	for j := p; j < r; j++ {
		if A[j] <= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]

	return i + 1
}

func printArrayWithPivot(A []int, p int, q int, r int) {
	for i := 0; i < len(A); i++ {
		if i == 0 && q == 0 {
			fmt.Printf(" [    ]")
			continue
		}

		switch i {
		case q:
			fmt.Printf(" ]")
		case 0, q + 1:
			fmt.Printf(" [")
		default:
		}

		if i < p || i > r {
			fmt.Printf("   ")
		} else {
			fmt.Printf(" %d", A[i])
		}
	}
	fmt.Printf(" ]")
}

func main() {
	A := []int{52, 16, 27, 37, 45, 68, 49, 11, 72, 38, 50}
	fmt.Println(A)
	// fmt.Println("A:", A, "p:", 0, "q:", 10)

	for i := 1; i < len(A); i++ {
		fmt.Println("===", i, "番目に小さい数 ===")
		fmt.Println(randomizedSelect(A, 0, len(A)-1, i))
		fmt.Println()
	}
}
