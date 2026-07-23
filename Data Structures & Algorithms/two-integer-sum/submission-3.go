/** 

#1 solution, brute force: 6min41s. 
Time: O(nˆ2), two fors operating through array
Space: O(1), helper slice, is the same size if n grows (0,1)


#2 solution
map, 33min27s
time: O(n) -> n size of array nums, 
space: O(n) -> map can be of size n too, it depends on the nums arr

**/
func twoSum(nums []int, target int) []int {

var twoSum []int
seen := make(map[int]int)

for i, num := range nums{
	x := target - num
	idxValue, ok := seen[x] 
	seen[num] = i

	if ok {
		twoSum = append(twoSum, idxValue, i)
		return twoSum
	} 
	}
return twoSum

}