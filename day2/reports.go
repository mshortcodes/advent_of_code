package main

type reports [][]int

// countSafe returns the count of safe reports.
func (r reports) countSafe() int {
	var count int

	for _, report := range r {
		if r.determineSafe(report) {
			count++
		}
	}

	return count
}

// determineSafe checks the conditions for a safe report.
func (r reports) determineSafe(report []int) bool {
	var isInc bool
	var isDec bool

	for i := 0; i < len(report)-1; i++ {
		nextNum := report[i+1]
		currentNum := report[i]

		if nextNum == currentNum {
			return false
		}
		if (nextNum < currentNum) && (nextNum >= currentNum-3) {
			isDec = true
			continue
		}
		if (nextNum < currentNum) && !(nextNum >= currentNum-3) {
			return false
		}
		if (nextNum > currentNum) && (nextNum <= currentNum+3) {
			isInc = true
			continue
		}
		if (nextNum > currentNum) && !(nextNum <= currentNum+3) {
			return false
		}
	}

	if isInc && isDec {
		return false
	} else if isDec {
		return true
	} else if isInc {
		return true
	} else {
		return false
	}
}
