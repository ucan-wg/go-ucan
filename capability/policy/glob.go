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

// globMatch matches a string against a pattern with '*' wildcards, handling escaped '\*' literals.
func globMatch(pattern, str string) bool {
	if !validateGlobPattern(pattern) {
		return false
	}

	// i is the index for the pattern
	// j is the index for the string
	var i, j int
	for i < len(pattern) && j < len(str) {
		switch pattern[i] {
		case '*':
			// skip consecutive '*' characters
			for i < len(pattern) && pattern[i] == '*' {
				i++
			}
			if i == len(pattern) {
				return true
			}

			// match the rest of the pattern
			for j < len(str) {
				if globMatch(pattern[i:], str[j:]) {
					return true
				}
				j++
			}

			return false
		case '\\':
			// handle escaped '*'
			i++
			if i < len(pattern) && pattern[i] == '*' {
				if str[j] != '*' {
					return false
				}
				i++
				j++
			} else {
				if i >= len(pattern) || pattern[i] != str[j] {
					return false
				}
				i++
				j++
			}
		default:
			if pattern[i] != str[j] {
				return false
			}
			i++
			j++
		}
	}

	// check for remaining characters in pattern
	for i < len(pattern) && pattern[i] == '*' {
		i++
	}

	return i == len(pattern) && j == len(str)
}
