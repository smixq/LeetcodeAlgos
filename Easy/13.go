package main

func romanToInt(s string) int {
	var sum int = 0
	var letHashMap = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i := range s {
		var add int
		add = letHashMap[rune(s[i])]
		if i != 0 {
            minus := letHashMap[rune(s[i-1])]
			if add > minus{
				sum = sum - (minus * 2)
			}
		}
		sum += add

	}
	return sum
}
