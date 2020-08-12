package main

import (
    "fmt"
)

func main(){
    nums := []int{2, 4, 1, 8, 5, 3, 0, 9, 7}
    quickSortC(nums)
    fmt.Println(nums)
    nums = []int{2, 4, 1, 8, 5, 3, 0, 9, 7}
    quickSortR(nums, 0, len(nums) - 1)
    fmt.Println(nums)
}

func partition(nums []int, low, high int)int{
    mid := nums[low]
    for low < high{
        for low < high && mid < nums[high] {
            high --
        }
        nums[low] = nums[high]
        for  low < high && mid > nums[low]{
            low ++
        }
        nums[high] = nums[low]
    }
    nums[low] = mid
    return low
}

func quickSortR(nums []int, left, right int){
    if len(nums) <= 1 {
        return
    }
    if left < right{
        mid := partition(nums, left, right)
        quickSortR(nums, left, mid -1)
        quickSortR(nums, mid+1, right)
    }

}
// 非递归版是借助了栈的特性，使用栈保存了不同分区的下标
func quickSortC(nums []int) {
    if len(nums) <= 1{
        return
    }
    mid := partition(nums, 0, len(nums) - 1)
    stack := make([]int, 0)
    if 0 < mid {
        stack = append(stack, 0, mid - 1)
    }
    if mid < len(nums) - 1{
        stack = append(stack, mid + 1, len(nums) - 1)
    }
    for len(stack) > 0{
        end := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        start := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        mid := partition(nums, start, end)
        if start < mid- 1 {
            stack = append(stack, start, mid - 1)
        }
        if end > mid + 1 {
            stack = append(stack, mid + 1, end)
        }
    }

    return
}
