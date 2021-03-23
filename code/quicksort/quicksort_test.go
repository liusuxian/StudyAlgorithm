package quicksort

import (
    "testing"
)

func QuickSort(nums []int) {
    QuickSortHelper(nums, 0, len(nums)-1)
}

func QuickSortHelper(nums []int, left, right int) {
    if left >= right {
        return
    }

    i, j, pivot := left, right, nums[left]
    for i < j {
        for i < j && nums[j] >= pivot {
            j--
        }
        nums[i] = nums[j]

        for i < j && nums[i] <= pivot {
            i++
        }
        nums[j] = nums[i]
    }

    nums[i] = pivot
    QuickSortHelper(nums, left, i-1)
    QuickSortHelper(nums, i+1, right)
}

func TestPermute(t *testing.T) {
    nums := []int{8, 3, 10, 2, 7, 6, 9, 12}
    t.Log(nums)
    QuickSort(nums)
    t.Log(nums)

    nums = []int{2, 2, 2, 1, 1, 3}
    t.Log(nums)
    QuickSort(nums)
    t.Log(nums)
}
