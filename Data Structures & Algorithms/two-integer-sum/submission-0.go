//#1 solution, brute force: 6min41s. O(nˆ2)

func twoSum(nums []int, target int) []int {

	var twoSum []int
    
	for i := 0; i <= len(nums)-1; i++ {
		for j:= i+1; j <= len(nums)-1; j++{
			if nums[i] + nums[j] == target {
				twoSum = append(twoSum, i, j)
			}
		}
	}
	return twoSum
}
