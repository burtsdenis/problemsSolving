package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{5,3,4,1,10,61,52}, 5))
}

func twoSum(nums []int, target int) []int {
    var sum []int
    for i := 0; i < len(nums); i++ {
        for j := 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target  && i != j {
               sum = append(sum, i, j)
                break
            }
        }
        if len(sum) == 2 { break }
    }
        return sum
}