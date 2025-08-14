package numbers

import (
	"math/rand"
	"slices"
)

type TplVarsNumbers struct {
	Game      string
	Target    []int
	Selection []int
}

var bigNumbers = []int{25, 50, 75, 100}
var smallNumbers = []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10}
var totalNumberCount = 6

func Pick(bigCount int) ([]int, []int) {
	smallCount := totalNumberCount - bigCount
	target := -1
	selection := []int{}
	big := slices.Clone(bigNumbers)
	small := slices.Clone(smallNumbers)

	for range bigCount {
		iRnd := rand.Intn(len(big))
		selection = append(selection, big[iRnd])
		big = slices.Delete(big, iRnd, iRnd+1)
	}

	for range smallCount {
		iRnd := rand.Intn(len(small))
		selection = append(selection, small[iRnd])
		small = slices.Delete(small, iRnd, iRnd+1)
	}

	for target <= 100 {
		target = rand.Intn(1000)
	}

	return splitInt(target), selection
}

func splitInt(n int) []int {
	if n == 0 {
		return []int{0}
	}
	tmp := n
	length := 0
	for tmp > 0 {
		tmp /= 10
		length++
	}
	digits := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		digits[i] = n % 10
		n /= 10
	}
	return digits
}
