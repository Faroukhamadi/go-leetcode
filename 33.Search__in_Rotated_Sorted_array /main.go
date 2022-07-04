package main

func main() {
}

func search(nums []int, target int) int {
	pivot := getPivotLocation(nums)
	l, r := 0, len(nums)-1

	switch {
	case pivot == -1:
	case target > nums[r]:
		r = pivot
	default:
		l = pivot
	}

	for l <= r {
		m := (l + r) / 2
		switch {
		case target == m:
			return m
		case target > nums[m]:
			l = m + 1
		default:
			r = m - 1
		}
	}
	return -1
}

// get pivot location
func getPivotLocation(nums []int) int {
	l, r := 0, len(nums)-1

	// non-rotated array
	if nums[l] < nums[r] {
		return -1
	}

	for l <= r {
		m := (l + r) / 2

		switch {
		case m < len(nums)-1 && nums[m] > nums[m+1]:
			return m
		case m > 0 && nums[m] < nums[m-1]:
			return m
		case nums[l] < nums[m]:
			l = m + 1
		default:
			r = m - 1
		}
	}
	return -1
}
