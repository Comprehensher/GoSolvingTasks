package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern, compileErr := regexp.Compile("[A-z]oat")
	description := "A boat for one person"
	question := "Is that a goat?"
	preference := "I like oats"
	if compileErr == nil {
		fmt.Println("Description:",
			pattern.MatchString(description))
		fmt.Println("Question:",
			pattern.MatchString(question))
		fmt.Println("Preference:",
			pattern.MatchString(preference))
	} else {
		fmt.Println("Error:", compileErr)
	}
}
