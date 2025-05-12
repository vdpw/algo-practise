package lc

import "fmt"

/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s 的，而不是部分字符串。

示例 1：
输入：s = "aa", p = "a"
输出：false
解释："a" 无法匹配 "aa" 整个字符串。

示例 2:
输入：s = "aa", p = "a*"
输出：true
解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。

示例 3：
输入：s = "ab", p = ".*"
输出：true
解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。

提示：
1 <= s.length <= 20
1 <= p.length <= 20
s 只包含从 a-z 的小写字母。
p 只包含从 a-z 的小写字母，以及字符 . 和 *。
保证每次出现字符 * 时，前面都匹配到有效的字符
*/

// https://leetcode.cn/problems/substring-with-concatenation-of-all-words/
func findSubstring(s string, words []string) []int {
	return nil
}

/*
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号
子串
的长度。

示例 1：
输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"
示例 2：

输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"
示例 3：

输入：s = ""
输出：0

提示：
0 <= s.length <= 3 * 10^4
s[i] 为 '(' 或 ')'
*/
func longestValidParentheses(s string) int {
	if len(s) <= 0 {
		return 0
	}
	l := len(s)
	if isValidParentheses(s[0:l]) {
		return len(s)
	}
	return max(longestValidParentheses(s[1:l]), longestValidParentheses(s[0:l-1]))
}

func isValidParentheses(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b == '(' || b == '[' || b == '{' {
			stack = append(stack, b)
			continue
		}
		if len(stack) > 0 {
			last := stack[len(stack)-1]
			if (b == ')' && last == '(') || (b == ']' && last == '[') || (b == '}' && last == '{') {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, b)
	}

	return len(stack) == 0
}

func Hanoi(n int, source, target, auxiliary string) {
	if n == 1 {
		fmt.Printf("Move %d from %s to %s\n", n, source, target)
		return
	}
	Hanoi(n-1, source, auxiliary, target)
	fmt.Printf("Move %d from %s to %s\n", n, source, target)
	Hanoi(n-1, auxiliary, target, source)
}
