package random

import (
	"math/rand"
	"time"
)

// Randomize возвращает случайное целое число на интервале [left, right]
func Randomize(left, right int) int {
	return left + rand.Intn(right-left)
}

// ShuffleSeed генерирует новое зерно для генерации случайных чисел
func ShuffleSeed() {
	rand.Seed(time.Now().Unix())
}

// Combination генерирует случайное сочетание из n по k
func Combination(n, k int) []int {
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
