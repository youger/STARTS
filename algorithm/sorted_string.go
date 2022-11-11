package main

func sorted(s string) []string {
	count := len(s)
	flags := make(map[byte]bool, count)
	var res []string = make([]string, 0)
	if count <= 1 {
		return []string{s}
	}
	for i := 0; i < count; i++ {
		first := s[i]
		if flags[first] {
			continue
		}
		flags[first] = true
		var left string
		if i == 0 {
			left = s[i+1:]
		} else if i == count-1 {
			left = s[:count-1]
		} else {
			left = s[:i] + s[i+1:]
		}
		sortedStrs := sorted(left)

		for _, val := range sortedStrs {
			res = append(res, string(first)+val)
		}
	}
	return res
}
