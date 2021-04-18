/*
 * @lc app=leetcode.cn id=299 lang=golang
 *
 * [299] 猜数字游戏
 */
//1.两层循环遍
// @lc code=start
func getHint(secret string, guess string) string {
	x, y := 0,0
	mapS := make(map[int32]int)
	for k, v := range secret {
		if v == int32(guess[k]){
			x++
		} 
		mapS[v]++
	}
	for _, v := range guess{
		if mapS[v] != 0{
			y++
			mapS[v]--
		} 
	}
	return fmt.Sprintf("%dA%dB", x, y-x)
}
// @lc code=end

