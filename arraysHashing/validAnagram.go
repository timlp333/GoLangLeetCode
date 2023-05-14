package arraysHashing

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	checkSMap := make(map[byte]int)
	checkTMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {

		if _, isExistS := checkSMap[s[i]]; !isExistS {
			checkSMap[s[i]] = 1
		} else {
			checkSMap[s[i]] += 1
		}
		if _, isExistT := checkTMap[t[i]]; !isExistT {
			checkTMap[t[i]] = 1
		} else {
			checkTMap[t[i]] += 1
		}
	}
	if len(checkSMap) != len(checkTMap) {
		return false
	}

	for key, element := range checkSMap {
		value, isExistTT := checkTMap[key]
		if !isExistTT || value != element {

			return false
		}
	}

	return true
}
