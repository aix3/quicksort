package quicksort

// QuickSort is a quick sort impl
func QuickSort(arr []int, left, right int) {
	if left == right {
		return
	}
	base := arr[left]
	var i = right
	var j = left

	for i > j {
		for arr[i] >= base && i > j {
			i--
		}
		for arr[j] <= base && j < i {
			j++
		}
		if i != j {
			swap(arr, i, j)
		}
	}
	if left != i {
		swap(arr, left, i)
	}

	if i > left {
		QuickSort(arr, left, i-1)
	}

	if j < right {
		QuickSort(arr, j+1, right)
	}
}

func swap(array []int, i, j int) {
	array[i] ^= array[j]
	array[j] ^= array[i]
	array[i] ^= array[j]
}
