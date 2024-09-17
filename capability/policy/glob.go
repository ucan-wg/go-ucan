package policy

// validateGlobPattern ensures the pattern conforms to the spec: only '*' and escaped '\*' are allowed.
func validateGlobPattern(pattern string) bool {
	for i := 0; i < len(pattern); i++ {
		if pattern[i] == '*' {
			continue
		}
		if pattern[i] == '\\' && i+1 < len(pattern) && pattern[i+1] == '*' {
			i++ // skip the escaped '*'
			continue
		}
		if pattern[i] == '\\' && i+1 < len(pattern) {
			i++ // skip the escaped character
			continue
		}
		if pattern[i] == '\\' {
			return false // invalid escape sequence
		}
	}

	return true
}

// globMatch matches a string against a pattern with * wildcards, handling escaped '\*' literals.
func globMatch(pattern, str string) bool {
	if !validateGlobPattern(pattern) {
		return false
	}

	// i is the index for the pattern
	// j is the index for the string
	var i, j int

	// starIdx keeps track of the position of the last * in the pattern.
	// matchIdx keeps track of the position in the string where the last * matched.
	var starIdx, matchIdx int = -1, -1

	for j < len(str) {
		if i < len(pattern) && (pattern[i] == str[j] || pattern[i] == '\\' && i+1 < len(pattern) && pattern[i+1] == str[j]) {
			// characters match or if there's an escaped character that matches
			if pattern[i] == '\\' {
				// skip the escape character
				i++
			}
			i++
			j++
		} else if i < len(pattern) && pattern[i] == '*' {
			// there's a * wildcard in the pattern
			starIdx = i
			matchIdx = j
			i++
		} else if starIdx != -1 {
			// there's a previous * wildcard, backtrack
			i = starIdx + 1
			matchIdx++
			j = matchIdx
		} else {
			// no match found
			return false
		}
	}

	// check for remaining characters in the pattern
	for i < len(pattern) && pattern[i] == '*' {
		i++
	}

	// the entire pattern is processed, it's a match
	return i == len(pattern)
}