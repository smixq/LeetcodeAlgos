package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reversed := 0
	coppyX := x
	for coppyX > 0 {
		reversed = reversed*10 + coppyX%10
		coppyX /= 10
	}
	return reversed == x
}
