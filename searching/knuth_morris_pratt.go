package searching

import (
	"io"
	"log"
)

func KmpSearch(text string, subString string) []int {
	enableLogging(false)
	patternStartIndicies := []int{}
	if len(text) < len(subString) || len(subString) == 0 {
		return patternStartIndicies
	}

	ssChars := []rune(subString)
	// build longest prefix suffix array to aide in skipping forward n places on mismatched substrings
	lps := createLps(ssChars)

	i, j := 0, 0
	chars := []rune(text)
	for i < len(chars) {
		// compare each character to each char in a substring
		if chars[i] == ssChars[j] {
			i++
			j++
		}
		if j == len(ssChars) {
			// IF we check all substring characters successfully, calculate and store the start index subtracting the length
			// of our substring from the current text string index
			// then move j forward _past_ any indicies part of the current match
			patternStartIndicies = append(patternStartIndicies, i-j)
			j = lps[j-1]
		} else if i < len(chars) && chars[i] != ssChars[j] {
			// OTHERWISE either skip j forward based on the LPS calculation, or simply increment the current text substring
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return patternStartIndicies
}

func createLps(subString []rune) []int {
	lps := make([]int, len(subString))
	j := 0

	// Preprocess the pattern (calculate lps array)
	for i := 1; i < len(subString); {
		log.Printf("substring[i=%v]: %v\n", i, string(subString[i]))
		log.Printf("substring[j=%v]: %v\n", j, string(subString[j]))
		if subString[i] == subString[j] {
			log.Printf("match, incrementing j\n")
			j++
			lps[i] = j
			i++
		} else {
			log.Printf("no match, ")
			if j != 0 {
				j = lps[j-1]
				log.Printf("j != 0, setting to lps[j-1] (%v) \n", j)
			} else {
				log.Printf("setting lps[%v] to zero \n", i)
				lps[i] = 0
				i++
			}
		}
		log.Printf("current lps: %v, i=%v, j=%v\n", lps, i, j)
	}

	return lps
}

func enableLogging(enable bool) {
	if enable {
		log.SetFlags(0)
	} else {
		log.SetOutput(io.Discard)
	}
}
