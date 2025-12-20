package balance

func Balance(s string) bool {
	stack := make([]rune, 0)

	for _, line := range s {
		if line == '{' || line == '[' || line == '(' {
			stack = append(stack, line)
		} else if line == '}' || line == ']' || line == ')' {
			if len(stack) == 0 {
				return false
			}

			last := stack[len(stack)-1]

			if (line == '}' && last != '{') ||
				(line == ']' && last != '[') ||
				(line == ')' && last != '(') {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
