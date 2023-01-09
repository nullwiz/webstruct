package strings

import (
	"fmt"
	"sort"
	"strings"

	"modernc.org/sortutil"
)

func RemoveDups(elements string) string {
	arr := strings.Split(elements, "")
	encountered := map[string]bool{}
	fmt.Println("elements:", arr)
	for v := range arr {
		if encountered[arr[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[arr[v]] = true
		}
	}
	// Create a new slice to hold the unique elements.
	var result []string
	for key := range encountered {
		result = append(result, key)
	}
	fmt.Println(result)
	// join array back into string
	return strings.Join(result, "")
}

// declare array function to check if palindrome
func IsPalindrome(s string) string {
	// convert string to rune array
	r := []rune(
		// remove spaces
		strings.Replace(
			// lowercase string
			strings.ToLower(s),
			// replace spaces with empty string
			" ", "", -1),
	)
	// reverse rune array
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	// if rune arraty is equal to original string, return true string
	if string(r) == string(s) {
		return "true"
	}
	return "false"

}

// declare array function to make string lowercase
func Lowercase(s string) string {
	return strings.ToLower(s)
}

// check if string is permutation of another string
func IsPermutation(s1, s2 string) bool {
	// convert strings to rune array
	r1 := []rune(s1)
	r2 := []rune(s2)
	// sort rune arrays
	sort.Sort(sort.Reverse(sortutil.RuneSlice(r1)))
	sort.Sort(sort.Reverse(sortutil.RuneSlice(r2)))
	// compare rune arrays
	return string(r1) == string(r2)
}

// check if string is a palindrome permutation of another string
func IsPalindromePermutation(s1, s2 string) bool {
	// convert strings to rune array
	r1 := []rune(s1)
	r2 := []rune(s2)
	// sort rune arrays
	sort.Sort(sort.Reverse(sortutil.RuneSlice(r1)))
	sort.Sort(sort.Reverse(sortutil.RuneSlice(r2)))
	// compare rune arrays
	return string(r1) == string(r2)
}
