/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */
/*
1. 有序数组：后一个和前一个进行比较，相同则删除
2. 快慢指针
*/
// @lc code=start
func removeDuplicates(nums []int) int {
	// for i := len(nums) - 1; i > 0; i-- {
	// 	if nums[i] == nums[i-1]{//删除
	// 		nums = append(nums[:i],nums[i+1:]...)
	// 	}
	// }
	// return len(nums)
	n := len(nums)
	if n < 2{
		return n
	}
	right, left := 1, 1
	for right < n{
		if nums[right] != nums[right-1]{
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}
// @lc code=end

