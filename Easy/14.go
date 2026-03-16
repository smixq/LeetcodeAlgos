package main

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    first := strs[0]
	for i := 0; i < len(first); i ++ {
        ch := first[i]
        for j := 1; j < len(strs); j ++ {
            if len(strs[j]) <= i  || ch != strs[j][i] {
                return first[:i]
            }
        }
    }
	return first
}

