
// force way, two loop
func twoSum(nums []int, target int) []int {
    var location []int
    for i := 0; i < len(nums); i ++ {
        for j := i + 1; j < len(nums); j ++ {
            x , y := nums[i], nums[j]
            if x + y == target {
                location = []int{i, j}
                break
            }
        } 
    }
    return location
}


//Way2: use array, not a good way, because you don't know the max vlue int original array.
func twoSumUseArray(nums []int, target int) []int {
    var newNums [1000]int
    for i:= 0; i < len(nums) - 1; i ++ {
        newNums[nums[i]] = i
    }
    var location []int
    var left int
    var right int
    for i:= 0; i < len(nums); i ++ {
        left = nums[i]
        right = target - left
        if right > 0 && newNums[right] >= 0 {
            location = []int{i, newNums[right]}
            break;
        }
    }
    return location
}

