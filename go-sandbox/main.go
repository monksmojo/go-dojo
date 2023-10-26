package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func main() {
	searchRange := func(numbers []int, target int) []int {
		res := []int{-1, -1}
		for i, v := range numbers {
			if v == target {
				if res[0] == -1 && res[1] == -1 {
					res[0] = i
					res[1] = i
				} else if i < res[0] {
					res[0] = i
				} else if i > res[1] {
					res[1] = i
				}
			} else if v > target {
				break
			}
		}
		return res
	}
	fmt.Println(searchRange([]int{5,7,7,8,8,10},8))
	fmt.Println(searchRange([]int{5,7,7,8,8,10},6))
	var ab int
	fmt.Scanln(&ab)

	N, M, K := 2, 3, 1
	MOD := int(1e9 + 7)
	DP := make([][][]int, N+1)

	for i := 0; i < len(DP); i++ {
		DP[i] = make([][]int, N+1)
		for j := 0; j < len(DP[i]); j++ {
			DP[i][j] = make([]int, M+1)
			for k := 0; k < len(DP[i][j]); k++ {
				DP[i][j][k] = -1
			}
		}
	}
	r3 := solve(0, 0, 0, &N, &M, &K, &DP, &MOD)
	fmt.Println(r3)

	// Print the slice to check the result

	consecutivePairOfAB := func(str string) (aCount int, bCount int) {
		for i := 1; i < len(str); i++ {
			if str[i-1] == byte('A') && str[i] == byte('A') {
				aCount += 1
			} else if str[i-1] == byte('B') && str[i] == byte('B') {
				bCount += 1
			}
		}
		return aCount, bCount
	}
	aCount, bCount := consecutivePairOfAB("AABBABAABB")
	fmt.Printf("consecutive pair of AB in %s is A Pair %d : B Pair %d \n", "AABBABAABB", aCount, bCount)

	aCount, bCount = consecutivePairOfAB("ABABAB")
	fmt.Printf("consecutive pair of AB in %s is A Pair %d : B Pair %d \n", "ABABAB", aCount, bCount)

	consecutiveTripletsOfAB := func(str string) (aCount int, bCount int) {
		for i := 1; i < len(str)-1; i++ {
			if str[i-1] == byte('A') && str[i] == byte('A') && str[i+1] == byte('A') {
				aCount += 1
			} else if str[i-1] == byte('B') && str[i] == byte('B') && str[i+1] == byte('B') {
				bCount += 1
			}
		}
		return aCount, bCount
	}

	aCount, bCount = consecutiveTripletsOfAB("ABABAB")
	fmt.Printf("consecutive triplets of AB in %s is A triplets %d : B triplets %d \n", "ABABAB", aCount, bCount)

	aCount, bCount = consecutiveTripletsOfAB("AAABBAAABBBBABAAA")
	fmt.Printf("consecutive triplets of AB in %s is A triplets %d : B triplets %d \n", "AAABBAAABBBBABAAA", aCount, bCount)

	consecutiveStreamOfAB := func(color string) (aCount, bCount int) {
		var frequencyMap = map[string]int{"A": 0, "B": 0}
		for _, v := range color {
			if v == 'A' {
				if frequencyMap["B"] >= 3 {
					bCount += 1
				}
				frequencyMap["B"] = 0
				frequencyMap["A"]++
			}
			if v == 'B' {
				if frequencyMap["A"] >= 3 {
					aCount += 1
				}
				frequencyMap["A"] = 0
				frequencyMap["B"]++
			}
		}
		if frequencyMap["A"] >= 3 {
			aCount += 1
		}
		if frequencyMap["B"] >= 3 {
			bCount += 1
		}
		return aCount, bCount
	}

	aCount, bCount = consecutiveStreamOfAB("AAABBAAABBBBABAAA")
	fmt.Printf("consecutive stream of AB in %s is A stream %d : B stream %d \n", "AAABBAAABBBBABAAA", aCount, bCount)

	winnerOfGameV1 := func(colors string) bool {
		var frequencyMap = map[string]int{"A": 0, "B": 0}
		var aChance int
		var bChance int
		for _, v := range colors {
			if string(v) == "A" {
				if frequencyMap["B"] >= 3 {
					bChance = bChance + (frequencyMap["B"] - 2)
				}
				frequencyMap["B"] = 0
				frequencyMap["A"]++
			} else if string(v) == "B" {
				if frequencyMap["A"] >= 3 {
					aChance = aChance + (frequencyMap["A"] - 2)
				}
				frequencyMap["A"] = 0
				frequencyMap["B"]++
			}
		}
		if frequencyMap["A"] >= 3 {
			aChance = aChance + (frequencyMap["A"] - 2)
		}

		if frequencyMap["B"] >= 3 {
			bChance = bChance + (frequencyMap["B"] - 2)
		}

		if aChance > bChance {
			return true
		} else if aChance > bChance {
			return false
		} else {
			return false // both are equal so alice lose
		}
	}

	b := winnerOfGameV1("ABBBBBBBAAA")
	fmt.Println(b)

	goodPairs := func(nums []int) (pairs int) {
		numLen := len(nums)
		for i := 0; i < numLen; i += 1 {
			for j := i + 1; j < numLen; j += 1 {
				if nums[i] == nums[j] {
					pairs += 1
				}
			}
		}
		return pairs
	}

	q := goodPairs([]int{1, 2, 3, 4})
	fmt.Println(q)

	majorityElements := func(numbers []int) []int {
		majority := len(numbers) / 3
		majoritySlice := make([]int, 0)
		frequencyMap := make(map[int]int)
		for _, v := range numbers {
			if !(frequencyMap[v] > majority) {
				frequencyMap[v]++
				if frequencyMap[v] > majority {
					majoritySlice = append(majoritySlice, v)
				}
			}
		}
		return majoritySlice
	}

	fmt.Println(majorityElements([]int{3, 2, 3}))

}

func solve(idx int, searchCost int, maxSOFar int, N *int, M *int, K *int, DP *[][][]int, MOD *int) int {
	if idx == *N {
		if searchCost == *K {
			return 1
		}
		return 0
	}

	if (*DP)[idx][searchCost][maxSOFar] != -1 {
		return (*DP)[idx][searchCost][maxSOFar]
	}

	var result int
	for i := 1; i <= *M; i++ {
		if i > maxSOFar {
			result = (result + solve(idx+1, searchCost+1, i, N, M, K, DP, MOD)) % *MOD
		} else {
			result = (result + solve(idx+1, searchCost, maxSOFar, N, M, K, DP, MOD)) % *MOD
		}

	}

	(*DP)[idx][searchCost][maxSOFar] = result % *MOD
	return result % *MOD
}

func winnerOfGameV2(colors string) bool {
	aMoves := 0
	bMoves := 0

	for i := 1; i < len(colors)-1; i++ {
		if colors[i] == 'A' {
			if colors[i-1] == 'A' && colors[i+1] == 'A' {
				aMoves++
			}
		} else {
			if colors[i-1] == 'B' && colors[i+1] == 'B' {
				bMoves++
			}
		}
	}
	return aMoves > bMoves && aMoves != 0
}

// WAP to find the factorial of a number.
func factorial(a int) int {
	fact := 1
	for a >= 2 {
		fact *= a
		a += -1
	}
	return fact
}

/*
DONE
11. Write a Python program to swap the consecutive characters in a string input by the user.
For example input: “jackiechan”, output: “ajkceihcna”
*/
// using built in function
func swapAdjacentCharacter(inputString string) string {
	strLength := len(inputString)
	var swapBytes = make([]byte, 0, strLength)
	fmt.Println()
	for i := 0; i < strLength-1; i = i + 2 {
		swapBytes = append(swapBytes, inputString[i+1], inputString[i])
	}
	if strLength%2 != 0 {
		swapBytes = append(swapBytes, inputString[strLength-1])
	}
	return string(swapBytes)
}

func swapAdjacentCharacterV2(inputString string) string {
	runeSlice := []rune(inputString)
	for i := 0; i < len(runeSlice)-1; i += 2 {
		runeSlice[i], runeSlice[i+1] = runeSlice[i+1], runeSlice[i]
	}
	return string(runeSlice)
}

/*
DONE
12. Write a Python program to input 2 strings i.e. s1 and s2 and find the presence of s1 in s2, if
present then replace it with equal number of # in it. Eg input: s1=”cd”
s2=”abcdefcdhjpcdert” , output: ab##ef##hjp##ert
*/
func encodeString(childString string, parentString string) string {
	hashString := strings.Repeat("#", len(childString))
	if strings.Contains(parentString, childString) {
		return strings.ReplaceAll(parentString, childString, hashString)
	}
	return fmt.Sprintf("%s does not exist in %s", childString, parentString)
}

/*
DONE
Write a Python program to remove the characters which have odd index values of a given
string
*/
func filterOdd(givenString string) string {
	evenCharacterSlice := []rune{}
	for index, value := range givenString {
		if index%2 == 0 {
			evenCharacterSlice = append(evenCharacterSlice, value)
		}
	}
	return string(evenCharacterSlice)
}

/*
DONE
To find frequency of each character of the string
*/
func charFrequency(gString string) map[rune]int {
	freqMap := make(map[rune]int)
	for _, v := range gString {
		freqMap[v]++
	}
	return freqMap
}

/*
DONE
wap to find the second largest element in the array
*/
func secondLargest(intSlice []int) int {
	l, sl := intSlice[0], math.MinInt
	for i := 1; i < len(intSlice)-1; i++ {
		value := intSlice[i]
		if value > l {
			l, sl = intSlice[i], l
		} else if value < l && value >= sl {
			sl = value
		}
	}
	return sl
}

/*
DONE
wap to find the second largest element in the array
*/
func secondLargestV2(intSlice []int) int {
	l, sl := intSlice[0], math.MinInt
	for i, j := 1, len(intSlice)-1; i < j; i, j = i+1, j-1 {
		var value int
		if intSlice[i] > intSlice[j] {
			value = intSlice[i]
		} else {
			value = intSlice[j]
		}

		if value > l {
			l, sl = intSlice[i], l
		} else if value < l && value >= sl {
			sl = value
		}
	}
	return sl
}

// 14. Write a Python function to convert a given string to all uppercase if it contains at least
// uppercase characters in the first 4 characters.
func toUpper(stringArg string) string {
	for i := 0; i < 4; i++ {
		if unicode.IsUpper(rune(stringArg[i])) {
			return strings.ToUpper(stringArg)
		}
	}
	return "no upperCase character found in first four character of the string"
}

// 15. Write a Python program to remove duplicate characters of a given string.
func removeDuplicate(inputString string) string {
	var uniqueChar strings.Builder
	charCheckMap := make(map[rune]bool)
	for _, char := range inputString {
		if !charCheckMap[char] {
			uniqueChar.WriteRune(char)
			charCheckMap[char] = true
		}
	}
	return uniqueChar.String()
}

func checkPalindrome(inputString string) bool {
	inputString = strings.ToLower(strings.ReplaceAll(inputString, " ", ""))
	for i, j := 0, len(inputString)-1; i <= j; i, j = i+1, j-1 {
		if inputString[i] != inputString[j] {
			return false
		}
	}
	return true
}

// 17. Write a python program to print its reverse.
func reverseTheString(inputString string) string {
	strLength := len(inputString)
	charSlice := make([]byte, strLength)
	for i := strLength - 1; i >= 0; i -= 1 {
		charSlice = append(charSlice, inputString[i])
	}
	return string(charSlice)
}

// 18. Write a Python program to get a single string from two given strings, separated by a space
// and swap the first two characters of each string. Eg input: 'abc', 'xyz' output: 'xyc abz'
func createString(str1, str2 string) string {
	var char1 byte = str1[0]
	var char2 byte = str2[1]
	str1 = str1[1:]
	str2 = str2[1:]
	return fmt.Sprintf("%s%s %s%s", string(char2), str1, string(char1), str2)
}

// 19. Write a Python program to add 'ing' at the end of a given string (length should be at least 3).
// If the given string already ends with 'ing' then add 'ly' instead. If the string length of the
// given string is less than 3, leave it unchanged. Eg input: 'abc' output: 'abcing' input: 'string'
// output: 'stringly'

func addPrefix(inputString string) string {
	strLength := len(inputString)
	const ing string = "ing"
	if strLength >= 3 {
		if inputString[strLength-3:strLength] == ing {
			return inputString + "ly"
		}

		return inputString + ing
	}
	return inputString
}

// 20. Write a Python function to reverses a string if it's length is a multiple of 4
func isMultipleOf4(inputString string) string {
	strLength := len(inputString)
	if strLength%4 == 0 {
		return reverseTheString(inputString)
	}
	return fmt.Sprintf("%v with length %d is not multiple of 4", inputString, strLength)
}

// 21. Write a Python program to take input slice and an upper value. All the values less than

func filterByUpperValue(inputSlice []int, upperValue int) (outputSlice []int) {
	for _, v := range inputSlice {
		if v > upperValue {
			outputSlice = append(outputSlice, v)
		}
	}
	return outputSlice
}

/*
We will not use built in function in filterByUpperValue2
*/
func filterByUpperValue2(inputSlice []int, upperValue int) []int {
	sliceLen := len(inputSlice)
	i, j := 0, sliceLen-1
	for i < j {
		if inputSlice[i] < upperValue {
			inputSlice[i], inputSlice[j] = inputSlice[j], inputSlice[i]
			j -= 1
			i -= 1
		}
		i += 1
	}
	return inputSlice[:j]
}

// inserting elements into the slice
func insert(slice []int, index, value int) ([]int, error) {
	fmt.Println(cap(slice))
	if len(slice) < index {
		return nil, fmt.Errorf("invalid index")
	} else if index == 0 {
		slice = append([]int{value}, slice...) // inserting element at the start
		return slice, nil
	} else if len(slice) == index {
		slice = append(slice, value)
		return slice, nil
	} else {
		slice = append(slice[:index+1], slice[index:]...)
		slice[index] = value
		return slice, nil
	}

}

// 23. Write a Python program to sum all the items in a list input by the user
func sumOfSlice(intSlice []int) (sum int) {
	for _, v := range intSlice {
		sum += v
	}
	return sum
}

func sumOfSliceV2(intSlice []int) (sum int) {
	for i, j := 0, len(intSlice)-1; i < j; i, j = i+1, j-1 {
		sum = sum + intSlice[i] + intSlice[j]
	}
	return sum
}

// 24. Write a Python program to get the largest number from a list.
func largestNum(nums []int) int {
	largest := nums[0]
	for i := 1; i <= len(nums)-1; i += 1 {
		if nums[i] > largest {
			largest = nums[i]
		}
	}
	return largest
}

func largestNumV2(nums []int) int {
	largest := nums[0]
	for i, j := 1, len(nums)-1; i < j; i, j = i+1, j-1 {
		if nums[i] > largest {
			largest = nums[i]
		}
		if nums[j] > largest {
			largest = nums[j]
		}
	}
	return largest
}

// 25. Write a Python program to find the second smallest number in a list.
func secondSmallest(nums []int) int {
	s, ss := nums[0], math.MinInt32
	for i := 1; i < len(nums); i += 1 {
		value := nums[i]
		if value < s {
			s, ss = value, s
		} else if value > s && value <= ss {
			ss = value
		}
	}
	return ss
}

// 26. Write a Python program to multiply all the items in a list input by the user that are odd and multiples of 5.
func getOddMultipleOf5(nums []int) (filterNums []int) {
	for _, v := range nums {
		if v%2 == 1 && v%5 == 0 {
			filterNums = append(filterNums, v)
		}
	}
	return filterNums
}

func getOddMultipleOf5V2(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if !(nums[i] == 1 && nums[i]%5 == 0) { // negation of odd and multiple of 5
			nums[i], nums[j] = nums[j], nums[i]
			i -= 1
			j -= 1
		}
	}
	return nums[:j]
}

/*
Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.
Return any array that satisfies this condition.
Example 1:
Input: nums = [3,1,2,4]
Output: [2,4,3,1]
Explanation: The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
*/

// using built in function
// this is slow and takes about
func getParity(nums []int) (parity []int) {
	for _, v := range nums {
		if v%2 == 0 {
			parity = append([]int{v}, parity...)
		} else {
			parity = append(parity, v)
		}
	}
	return parity
}

// without using a built in function
func getParity1(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]%2 != 0 && nums[j]%2 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i += 1
			j -= 1
		} else if nums[i]%2 == 0 {
			i = +1
		} else if nums[j]%2 != 0 {
			j -= 1
		}
	}
	return nums
}

func isMonotonic(nums []int) bool {
	i, j := 0, 1
	return monotonicDirection(nums, i, j)
}

func monotonicDirection(nums []int, i, j int) bool {
	if j < len(nums) && nums[i] == nums[j] {
		i += 1
		j += 1
		return monotonicDirection(nums, i, j)
	} else if j < len(nums) && nums[j] > nums[i] {
		i += 1
		j += 1
		return isMonotonicIncreasing(nums, i, j)
	} else {
		i += 1
		j += 1
		return isMonotonicDecreasing(nums, i, j)
	}
}

func isMonotonicIncreasing(nums []int, i, j int) bool {
	for j < len(nums) {
		if !(nums[j] >= nums[i]) {
			return false
		}
		i += 1
		j += 1
	}
	return true
}

func isMonotonicDecreasing(nums []int, i, j int) bool {
	for j < len(nums) {
		if !(nums[i] >= nums[j]) {
			return false
		}
		i += 1
		j += 1
	}
	return true
}

func isMonotonic2(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	direction := 0
	for i := 0; i < len(nums)-1; i += 1 {
		if nums[i] < nums[i+1] {
			if direction == 0 {
				direction = 1
			} else if direction == -1 {
				return false
			}
		}

		if nums[i] > nums[i+1] {
			if direction == 0 {
				direction = -1
			} else if direction == 1 {
				return false
			}
		}
	}
	return true
}

/*

456. 132 Pattern
Given an array of n integers nums, a 132 pattern is a subsequence of three integers nums[i], nums[j] and nums[k] such that i < j < k and nums[i] < nums[k] < nums[j].

Return true if there is a 132 pattern in nums, otherwise, return false.

Example 1:

Input: nums = [1,2,3,4]
Output: false
Explanation: There is no 132 pattern in the sequence.
Example 2:

Input: nums = [3,1,4,2]
Output: true
Explanation: There is a 132 pattern in the sequence: [1, 4, 2].

*/

func check123Patter(nums []int) bool {
	third := math.MinInt32
	stack := []int{}
	for i := len(nums) - 1; i >= 0; i -= 1 {
		if nums[i] < third {
			return true
		}

		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			third = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}

/*
Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

Example 1:

Input: s = "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"
Example 2:

Input: s = "God Ding"
Output: "doG gniD"
*/
func reverseWords(sentence string) string {
	var rev = make([]rune, 0, len(sentence))
	wordsSlice := strings.Fields(sentence)
	for _, word := range wordsSlice {
		for i := len(word) - 1; i >= 0; i -= 1 {
			rev = append(rev, rune(word[i]))
		}
		rev = append(rev, ' ')
	}
	return string(rev)[:len(rev)-1]
}

func reverseWordsV2(sentence string) string {
	var stack = make([]rune, 0)
	var rev = make([]rune, 0, len(sentence))
	for _, v := range sentence {
		if v == ' ' {
			rev = popAll(&stack, rev)
			rev = append(rev, ' ')
		} else {
			stack = append(stack, v)
		}
	}
	rev = popAll(&stack, rev)
	return string(rev)
}

func popAll(stack *[]rune, rev []rune) []rune {
	for !isEmpty(stack) {
		top := pop(stack)
		rev = append(rev, top)
	}
	return rev
}

func isEmpty(stack *[]rune) bool {
	return len(*stack) == 0
}

func pop(stack *[]rune) rune {
	stack_size := len(*stack) - 1
	top := (*stack)[stack_size]
	*stack = (*stack)[:stack_size]
	return top
}

func swap(a, b *int) {
	fmt.Println(a, b)
	fmt.Println(*a, *b)
}
