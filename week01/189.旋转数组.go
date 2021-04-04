/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */
/*
*/
// @lc code=start
func rotate(nums []int, k int)  {
	inverse(nums)
	inverse(nums[:k%len(nums)])
	inverse(nums[k%len(nums):])
}
func inverse(arr []int){
	for i := 0; i < len(arr)/2; i++{
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
// @lc code=end

