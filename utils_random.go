package main

import (
	"math/rand"
	"time"
)

// randomize возвращает случайное целое число на интервале [left, right]
func randomize(left, right int) int {
	return left + rand.Intn(right-left)
}

// shuffleSeed генерирует новое зерно для генерации случайных чисел
func shuffleSeed() {
	rand.Seed(time.Now().Unix())
}

// randomCombination генерирует случайное сочетание из n по k
func randomCombination(n, k int) []int {
	set := make([]int, n)

	for i := 0; i < n; i++ {
		if i < k {
			set[i] = 1
		} else {
			set[i] = 0
		}
	}

	rand.Shuffle(n, func(i, j int) {
		set[i], set[j] = set[j], set[i]
	})

	ans := make([]int, k)

	for i, counter := 0, 0; i < n; i++ {
		if set[i] == 1 {
			ans[counter] = i
			counter++
		}
	}

	return ans
}
