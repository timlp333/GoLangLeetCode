package arraysHashing

// strs ex: ["eat","tea","tan","ate","nat","bat"]
func GroupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	// 利用array [26]int 26字母計算表 當map的key，方便歸類
	resultMap := make(map[[26]int][]string)

	for _, value := range strs {
		var check [26]int
		for i := 0; i < len(value); i++ {
			// a in memory is 97 , 97-97 = 0, 所以第一個字是a的話，check[97-97]+1
			// b 會在check[1]做計算
			check[value[i]-'a']++
		}
		// 將相同結果的字歸類 "bat" -> a1b1t1, ate -> a1e1t1 , eat -> a1e1t1
		// ate, eat 同類, bat自己一類
		if _, isExist := resultMap[check]; !isExist {
			resultMap[check] = []string{value}
		} else {
			resultMap[check] = append(resultMap[check], value)
		}
	}
	for _, value := range resultMap {
		ans = append(ans, value)
	}

	return ans
}
