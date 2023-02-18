package main

/*
给定一个字符串str，只由‘X’和‘.’两种字符构成。
‘X’表示墙，不能放灯，也不需要点亮
‘.’表示居民点，可以放灯，需要点亮
如果灯放在i位置，可以让i-1，i和i+1三个位置被点亮
返回如果点亮str中所有需要点亮的位置，至少需要几盏灯
*/
func minLight2(road string) int {
	str := []byte(road)
	light := 0
	for i := 0; i < len(str); {
		if str[i] == 'x' {
			i++
		} else {
			light++
			if i+1 == len(str) {
				break
			} else {
				if str[i+1] == 'x' {
					i = i + 2
				} else {
					i = i + 3
				}
			}
		}
	}
	return light
}

func main() {

}
