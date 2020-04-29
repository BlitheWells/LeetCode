
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

