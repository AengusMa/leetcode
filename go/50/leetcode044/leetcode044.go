package leetcode044

func isMatch(s string, p string) bool {

	return help([]byte(s), []byte(p))
}

func help(s, p []byte) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	if p[0] == '*' {

	} else if p[0] == '?' {

	} else {

	}

	return false
}
