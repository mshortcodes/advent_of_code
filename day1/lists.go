package main

type lists struct {
	left  []int
	right []int
}

// getDistance calculates the sum of the differences (absolute) between each list's numbers.
func (l lists) getDistance() int {
	var distance int

	for i := 0; i < len(l.left); i++ {
		switch {
		case l.left[i] < l.right[i]:
			distance += l.right[i] - l.left[i]
		case l.left[i] > l.right[i]:
			distance += l.left[i] - l.right[i]
		default:
		}
	}

	return distance
}

// getSimilarity calculates the similarity score by
// accumulating the product of each number in the left list
// its frequency in the right list.
func (l lists) getSimilarity() int {
	var similarity int

	for _, leftNum := range l.left {
		var count int
		for _, rightNum := range l.right {
			if leftNum == rightNum {
				count++
			}
		}
		similarity += leftNum * count
	}

	return similarity
}
