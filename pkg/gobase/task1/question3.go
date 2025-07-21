package task1

import "container/list"

// 有效的括号
func IsValid(s string) bool {
	l := list.New()

	for _, v := range s {
		// 入栈
		if v == '(' || v == '[' || v == '{' {
			l.PushBack(v)
		} else {
			// 出栈
			if l.Len() == 0 {
				return false
			}
			last := l.Back()
			if last.Value.(rune) == '(' && v == ')' || last.Value.(rune) == '[' && v == ']' || last.Value.(rune) == '{' && v == '}' {
				l.Remove(last)
			} else {
				return false
			}
		}
	}
	return l.Len() == 0
}
