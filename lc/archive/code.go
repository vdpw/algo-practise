package archive

// https://leetcode.cn/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}

// https://leetcode.cn/problems/add-two-numbers/description/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	cur := root
	third := 0
	for l1 != nil || l2 != nil || third != 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + third
		third = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	return root.Next
}

// https://leetcode.cn/problems/string-to-integer-atoi/
func myAtoi(s string) int {
	num := 0
	sign := 1
	start := false
	const max = 1<<31 - 1
	for _, c := range s {
		if c == ' ' {
			if start {
				break
			}
			continue
		}
		if c == '-' {
			if start {
				break
			}
			start = true
			sign = -1
			continue
		}
		if c == '+' {
			if start {
				break
			}
			start = true
			continue
		}
		if c < '0' || c > '9' {
			break
		}
		start = true

		digit := int(c - '0')
		if num > max/10 || (num == max/10 && digit > 7) {
			if sign == 1 {
				return max
			}
			return -1 << 31
		}
		num = num*10 + digit
	}
	return num * sign
}

// https://leetcode.cn/problems/zigzag-conversion/description/
func convert(s string, numRows int) string {
	totalLen := len(s)
	if numRows == 1 || totalLen <= numRows {
		return s
	}

	resultIdx := 0
	result := make([]byte, totalLen)

	for rowNum := 0; rowNum < numRows; rowNum++ {
		d := (numRows << 1) - 2
		a01 := rowNum
		a02 := d - rowNum
		for {
			if a01 >= totalLen {
				break
			}

			result[resultIdx] = s[a01]
			resultIdx++
			a01 += d

			if rowNum > 0 && rowNum < numRows-1 && a02 < totalLen {
				result[resultIdx] = s[a02]
				resultIdx++
				a02 += d
			}

		}
	}
	return string(result)
}

// https://leetcode.cn/problems/reverse-integer/description/
func reverse(x int) int {
	const MIN = (-1 << 31) / 10
	const MAX = ((1 << 31) - 1) / 10
	var res int

	for x != 0 {
		digit := x % 10
		if res > MAX || (res == MAX && digit > 7) {
			return 0
		}
		if res < MIN || (res == MIN && digit < -8) {
			return 0
		}
		res = res*10 + x%10
		x /= 10
	}

	return res
}

// https://leetcode.cn/problems/palindrome-number/description/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	origin := x
	const MAX = ((1 << 31) - 1) / 10
	reverse := 0

	for x > 0 {
		if reverse > MAX {
			return false
		}
		reverse = reverse*10 + x%10
		x /= 10
	}
	return reverse == origin
}

// https://leetcode.cn/problems/container-with-most-water/
func maxArea(height []int) int {

	max := 0
	for i, j := 0, len(height)-1; i < j; {
		area := min(height[i], height[j]) * (j - i)
		if area > max {
			max = area
		}
		// 在长度缩小的情况下，如果高度不变或者变小，那么面积一定变小
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

// https://leetcode.cn/problems/integer-to-roman/
func intToRoman(num int) string {
	mapping := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	result := ""
	for {
		if num == 0 {
			break
		}
		if num >= 1000 {
			result += mapping[1000]
			num -= 1000
		} else if num >= 900 {
			result += mapping[900]
			num -= 900
		} else if num >= 500 {
			result += mapping[500]
			num -= 500
		} else if num >= 400 {
			result += mapping[400]
			num -= 400
		} else if num >= 100 {
			result += mapping[100]
			num -= 100
		} else if num >= 90 {
			result += mapping[90]
			num -= 90
		} else if num >= 50 {
			result += mapping[50]
			num -= 50
		} else if num >= 40 {
			result += mapping[40]
			num -= 40
		} else if num >= 10 {
			result += mapping[10]
			num -= 10
		} else if num >= 9 {
			result += mapping[9]
			num -= 9
		} else if num >= 5 {
			result += mapping[5]
			num -= 5
		} else if num >= 4 {
			result += mapping[4]
			num -= 4
		} else {
			result += mapping[1]
			num--
		}
	}

	return result
}

// https://leetcode.cn/problems/roman-to-integer/
func romanToInt(s string) int {
	mapping := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	result := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if v, ok := mapping[s[i:i+2]]; ok {
				result += v
				i++
				continue
			}
		}
		result += mapping[s[i:i+1]]
	}
	return result
}

// https://leetcode.cn/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	minLen := len(strs[0])
	for _, str := range strs[1:] {
		if len(str) < minLen {
			minLen = len(str)
		}
	}
	if minLen == 0 {
		return ""
	}

	for i := 0; i < minLen; i++ {
		c := strs[0][i]
		for _, str := range strs[1:] {
			if str[i] != c {
				return str[:i]
			}
		}
	}
	return strs[0][:minLen]
}

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	mapping := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	result := []string{""}
	for i := 0; i < len(digits); i++ {
		digit := digits[i]
		letters, ok := mapping[digit]
		if !ok {
			continue
		}
		temp := make([]string, 0, len(result)*len(letters))
		for _, prefix := range result {
			for j := 0; j < len(letters); j++ {
				temp = append(temp, prefix+string(letters[j]))
			}
		}
		result = temp
	}
	return result
}

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// 将 fast 剩余节点扫描完毕, 即表示 slow 到达倒数第 n 个节点
	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

// https://leetcode.cn/problems/valid-parentheses/
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

// https://leetcode.cn/problems/merge-two-sorted-lists/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	cursor := result
	for {
		if list1 == nil && list2 == nil {
			break
		}
		if list1 == nil {
			cursor.Next = list2
			break
		}
		if list2 == nil {
			cursor.Next = list1
			break
		}
		if list1.Val < list2.Val {
			cursor.Next = list1
			list1 = list1.Next
		} else {
			cursor.Next = list2
			list2 = list2.Next
		}
		cursor = cursor.Next
	}
	return result.Next
}

// https://leetcode.cn/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	var result []string
	generateParenthesis_backtrack("", 0, 0, n, &result)
	return result
}

func generateParenthesis_backtrack(s string, open, close, max int, result *[]string) {
	if len(s) == max*2 {
		*result = append(*result, s)
		return
	}
	if open < max {
		generateParenthesis_backtrack(s+"(", open+1, close, max, result)
	}
	if close < open {
		generateParenthesis_backtrack(s+")", open, close+1, max, result)
	}
}

// https://leetcode.cn/problems/merge-k-sorted-lists/description/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	results := lists
	for {
		temp := make([]*ListNode, 0, len(results)/2+1)
		for i := 0; i < len(results); i += 2 {
			if i+1 < len(results) {
				temp = append(temp, mergeTwoLists(results[i], results[i+1]))
			} else {
				temp = append(temp, results[i])
			}
		}
		if len(temp) == 1 {
			return temp[0]
		}
		results = temp
	}
}

// https://leetcode.cn/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	return reverseKGroup(head, 2)
}

// https://leetcode.cn/problems/reverse-nodes-in-k-group/
func reverseKGroup(head *ListNode, k int) *ListNode {
	curr := head
	count := 0

	// Check if there are at least k nodes left in the linked list
	for curr != nil && count < k {
		curr = curr.Next
		count++
	}

	// If we have k nodes, then we reverse them
	if count == k {
		// Reverse first k nodes
		reversedHead := reverseKNodes(head, k)
		// head is now the end of the reversed group, connect it with the result of next reversal
		head.Next = reverseKGroup(curr, k)
		return reversedHead
	} else {
		// Less than k nodes, return head as is
		return head
	}
}

// Helper function to reverse k nodes
func reverseKNodes(head *ListNode, k int) *ListNode {
	var prev *ListNode = nil
	curr := head
	next := (*ListNode)(nil)
	count := 0

	// Reverse k nodes
	for count < k {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		count++
	}
	// prev is the new head of the reversed list
	return prev
}

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

// https://leetcode.cn/problems/remove-element/
func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return right
}

// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == len(needle)-1 {
				return i
			}
		}
	}
	return -1
}

// https://leetcode.cn/problems/search-in-rotated-sorted-array/
func search_in_rotated_sorted_array(nums []int, target int) int {
	reverseIdx := search_in_rotated_sorted_array_findReversePoint(nums, 0, len(nums)-1)
	if target == nums[reverseIdx] {
		return reverseIdx
	}
	if target > nums[reverseIdx] {
		return -1
	}
	if target >= nums[0] {
		return search_in_rotated_sorted_array_biSearch(nums, 0, reverseIdx, target)
	} else {
		return search_in_rotated_sorted_array_biSearch(nums, reverseIdx+1, len(nums)-1, target)
	}
}

func search_in_rotated_sorted_array_findReversePoint(nums []int, startIdx, endIdx int) (index int) {
	if startIdx == endIdx {
		return startIdx
	}
	midIdx := startIdx + (endIdx-startIdx)/2
	midVal := nums[midIdx]
	if midVal > nums[midIdx+1] {
		return midIdx
	}
	if nums[midIdx] < nums[startIdx] {
		return search_in_rotated_sorted_array_findReversePoint(nums, startIdx, midIdx)
	} else {
		return search_in_rotated_sorted_array_findReversePoint(nums, midIdx+1, endIdx)
	}
}

func search_in_rotated_sorted_array_biSearch(nums []int, startIdx, endIdx, target int) (index int) {
	if startIdx > endIdx {
		return -1
	}
	midIdx := startIdx + (endIdx-startIdx)/2
	if nums[midIdx] == target {
		return midIdx
	}
	if nums[midIdx] < target {
		return search_in_rotated_sorted_array_biSearch(nums, midIdx+1, endIdx, target)
	} else {
		return search_in_rotated_sorted_array_biSearch(nums, startIdx, midIdx-1, target)
	}
}

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	first := searchRange_lowerBound(nums, target)
	if first == len(nums) || nums[first] != target {
		return []int{-1, -1}
	}
	last := searchRange_upperBound(nums, target) - 1
	return []int{first, last}
}

func searchRange_lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func searchRange_upperBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

/*
https://leetcode.cn/problems/count-and-say
func countAndSay(n int) string {
	for {
		if n == 1 {
			return "1"
		}
		if n == 2 {
			return "11"
		}
		s := countAndSay(n - 1)
		var res []byte
		for i, j := 0, 0; i < len(s); i = j {
			for j < len(s) && s[j] == s[i] {
				j++
			}
			res = append(res, byte(j-i)+'0', s[i])
		}
		return string(res)
	}
}

func countAndSay(n int) string {
    pre := "1"
    for i:=2; i<=n; i++ {
        cur := []byte{}
        for j, start := 0, 0; j<len(pre); start = j {
            for j<len(pre) && pre[j] == pre[start] {
                j++
            }
            cur = append(cur, []byte(strconv.Itoa(j-start))...)
            cur = append(cur, pre[start])
        }
        pre = string(cur)
    }
    return pre
}

var preCalculatedData = map[int]string{ 1: "1", 2: "11", 3: "21" }
var arr = []string{1, 11, 21, 1211}
*/

// https://leetcode.cn/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
func LengthOfLongestSubstring(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}
	left := 0
	m := 0

	set := make(map[byte]int)
	for cursor := 0; cursor < length; cursor++ {
		b := s[cursor]
		if idx, found := set[b]; found && left <= idx {
			left = idx + 1
		}
		m = max(cursor-left+1, m)
		set[b] = cursor
	}
	m = max(length-left, m)
	return m
}

// https://leetcode.cn/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	// another O(n) method: https://cp-algorithms.com/string/manacher.html
	if len(s) < 2 {
		return s
	}

	// indices of the best window so far
	start, end := 0, 0

	// helper that grows outward while s[l]==s[r]
	expand := func(l, r int) (int, int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		return l + 1, r - 1 // bounds after the last match
	}

	for i := 0; i < len(s); i++ {
		// odd‑length palindrome (center on one rune)
		l1, r1 := expand(i, i)
		// even‑length palindrome (center between two runes)
		l2, r2 := expand(i, i+1)

		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}

	return s[start : end+1]
}

// https://leetcode.cn/problems/regular-expression-matching/
func isMatch(s string, p string) bool {
	// todo: learn dp
	m, n := len(s), len(p)

	// dp[i][j] == does s[i:] match p[j:] ?
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[m][n] = true // two empty strings match

	// iterate backwards so sub‑problems to the right / below are ready
	for i := m; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// firstCharMatch: s[i] exists && (equal or pattern has '.')
			firstCharMatch := i < m && (p[j] == s[i] || p[j] == '.')

			if j+1 < n && p[j+1] == '*' {
				// Case 1: '*' uses **zero** of the preceding element → skip "X*"
				// Case 2: '*' uses ≥1  → only if first char matches, stay at j
				dp[i][j] = dp[i][j+2] || (firstCharMatch && dp[i+1][j])
			} else {
				// simple char / '.' : both advance one step
				dp[i][j] = firstCharMatch && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}

// https://leetcode.cn/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for _, str := range strs {
		counter := [26]int{}
		for _, c := range str {
			counter[c-'a']++
		}
		m[counter] = append(m[counter], str)
	}
	res := [][]string{}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// https://leetcode.cn/problems/single-number/
func singleNumber(nums []int) int {
	a := 0
	for _, b := range nums {
		a ^= b
	}
	return a
}

// https://leetcode.cn/problems/add-binary/
func addBinary(a string, b string) string {
	la, lb := len(a), len(b)
	var dc byte = 0
	res := ""
	for {
		la--
		lb--
		var da, db byte = 0, 0
		co := false
		if la >= 0 {
			da = a[la] - '0'
			co = true
		}
		if lb >= 0 {
			db = b[lb] - '0'
			co = true
		}
		c := da + db + dc
		if c == 3 {
			dc = 1
			res = "1" + res
		} else if c == 2 {
			res = "0" + res
			dc = 1
		} else if c == 1 {
			res = "1" + res
			dc = 0
		} else {
			if co {
				res = "0" + res
			}
		}
		if !co {
			return res
		}
	}
}

// https://leetcode.cn/problems/plus-one/
func plusOne(digits []int) []int {
	b := true
	for r := len(digits) - 1; r >= 0; r-- {
		if digits[r] == 9 {
			digits[r] = 0
			continue
		} else {
			digits[r]++
			b = false
			break
		}
	}
	if b {
		return append([]int{1}, digits...)
	}
	return digits
}

// https://leetcode.cn/problems/sqrtx/
func mySqrt(x int) int {
	var table = [65536]int{}
	for i := 0; i < 65536; i++ {
		table[i] = i * i
	}
	l, r := 0, len(table)-1
	for {
		if l == r || l == r-1 {
			return l
		}
		m := (l + r) / 2
		v := table[m]
		if v > x {
			r = m
		} else if v < x {
			l = m
		} else {
			return m
		}
	}
}

// https://leetcode.cn/problems/valid-perfect-square/
func isPerfectSquare(num int) bool {
	var table = map[int]struct{}{}
	for i := 0; i < 65536; i++ {
		table[i*i] = struct{}{}
	}
	_, found := table[num]
	return found
}
