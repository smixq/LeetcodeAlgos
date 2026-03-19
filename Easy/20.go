package main

func isValid(s string) bool {
	var stacBrack []rune

	for _, brack := range s {
		n := len(stacBrack)
		if brack == '(' || brack == '[' || brack == '{' {
			stacBrack = append(stacBrack, brack)
		} else if n == 0 {
			return false
		} else {

			if (brack == ')' && stacBrack[n-1] == '(') || (brack == ']' && stacBrack[n-1] == '[') || (brack == '}' && stacBrack[n-1] == '{') {
				stacBrack = append(stacBrack[:len(stacBrack)-1], stacBrack[len(stacBrack):]...)
			} else {
				return false
			}
		}
	}
	return len(stacBrack) == 0
}
