/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	if len(nums) == 1{
		return [][]int{nums}
	}
	res := [][]int{}

	for i, num := range nums {
		temp := make([]int, len(nums)-1)
		copy(temp[0:], nums[0:i])
		copy(temp[i:], nums[i+1:])

		sub := permute(temp)
		for _, s := range sub {
			res = append(res, append(s, num))
		}
	}
	return res
}
// @lc code=end
