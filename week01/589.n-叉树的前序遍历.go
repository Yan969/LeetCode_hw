/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N 叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    if root == nil{
		return []int{}
	}
	var ans []int
	ans = append(ans, root.Val)
	for _, v := range root.Children{
		ans = append(ans, preorder(v)...)
	}
	return ans
}
// @lc code=end

