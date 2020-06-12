package main

import (
	"fmt"
	"regexp"
)

func main() {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}
	len := len(input)

	//조건 길이 패턴 + 길이 10이상
	pattern1 := "[0-9]"
	pattern2 := "[a-z]"
	pattern3 := "[A-Z]"
	pettern4 := "[~!@#$%^&*()_+|<>?:{}]"

	level1 := "LEVEL1"
	level2 := "LEVEL2"
	level3 := "LEVEL3"
	level4 := "LEVEL4"
	level5 := "LEVEL5"

	patternNum, _ := regexp.MatchString(pattern1, input)
	patternLower, _ := regexp.MatchString(pattern2, input)
	patternUppder, _ := regexp.MatchString(pattern3, input)
	patternSpe, _ := regexp.MatchString(pettern4, input)

	if patternNum && !patternLower && !patternUppder && !patternSpe {
		fmt.Println(level1)
	} else if !patternNum && patternLower && !patternUppder && !patternSpe {
		fmt.Println(level2)
	} else if patternNum && patternLower && patternUppder && !patternSpe {
		fmt.Println(level3)
	} else if patternNum && patternLower && patternUppder && patternSpe && len <= 10 {
		fmt.Println(level4)
	} else if patternNum && patternLower && patternUppder && patternSpe && len > 10 {
		fmt.Println(level5)
	}
}
