// Last updated: 24/07/2025, 12:49:53
// @param nums: array of integers, can be negative
// @param target: sum target, can be negative
func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    
    for i, val := range nums {
        j, ok := seen[target - val]
        if ok {
            return []int{j, i}
        }
        seen[val] = i
    }

    return nil
}