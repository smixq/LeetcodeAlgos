package main

func generateParenthesis(n int) []string {
	return filterBacks(generateAllBacks(n), n)
}

func generateAllBacks(n int) map[string]struct{} {
	n = n * 2
	res := make(map[string]struct{})
	res[")"] = struct{}{}
	res["("] = struct{}{}
	for i := 1; i < n; i++ {
		next := make(map[string]struct{})
		for key := range res {
			next[key+"("] = struct{}{}
			next[key+")"] = struct{}{}
		}
		res = next
	}

	return res
}
func filterBacks(backsMap map[string]struct{}, n int) []string {
	n = n * 2
	var res []string
	for key := range backsMap {
		openBacks := 0

		for _, back := range key {
			if back == '(' {
				openBacks += 1
			} else {
				if openBacks > 0 {
					openBacks -= 1
				} else {
					openBacks = -1
					break
				}
			}
		}
		if openBacks == 0 {
			res = append(res, key)
		}

	}
	return res
}

// ()()()
