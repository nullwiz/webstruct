package array

import (
	"fmt"
	"log"
	"sort"
	"strings"
	d "webstruct/database"

	"modernc.org/sortutil"
)

type ArrayInterface interface {
	opHandler()
}

type ArrayMethods interface {
	RemoveDuplicates()
	RemoveDuplicatesFromStringLiteral()
	IsPalindrome()
	Lowercase()
	IsPalindromePermutation()
	IsPermutation()
}

type ArrayH struct {
	handler  d.ArrayHandler
	database d.StructDB
}

func (c ArrayH) validateOp() error {
	var arrayOps = []string{"add", "remove", "get", "remDups", "toLowcase", "isPalindrome"}
	for _, opVal := range arrayOps {
		if opVal == c.handler.Operation {
			return nil
		}
	}
	return fmt.Errorf("invalid operation '%s'", c.handler.Operation)
}

func (c ArrayH) opHandler() (int, error) {
	if c.validateOp() == nil {
		switch c.handler.Operation {
		case "add":
			_, err := c.database.AddArray(c.handler, c.handler)
			if err != nil {
				log.Fatal("error when adding array")
			}
			return 1, nil
		case "get":
			_, err := c.database.GetArray(c.handler, c.handler)
			if err != nil {
				log.Fatal("error when getting array")
			}
			return 1, nil
		}
	}
	return 0, fmt.Errorf("invalid operation")
}

func RemoveDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	fmt.Println("elements:", elements)
	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
		}
	}
	// Create a new slice to hold the unique elements.
	var result []string
	for key := range encountered {
		result = append(result, key)
	}
	fmt.Println(result)
	return result
}

// remove duplicates from string literal
func RemoveDuplicatesFromStringLiteral(s string) string {
	// split string into array
	arr := strings.Split(s, "")
	// remove duplicates from array
	arr = RemoveDuplicates(arr)
	// join array back into string
	return strings.Join(arr, "")
}

// declare array function to check if palindrome
func IsPalindrome(s string) bool {
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
	return string(r) == string(s)
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
