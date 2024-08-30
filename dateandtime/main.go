package main

import (
	// "fmt"
	"fmt"
	"time"
)

func main() {
	d, err := time.ParseDuration("1h30m")
	if err == nil {
		Printfln("Hours: %v", d.Hours())
		Printfln("Mins: %v", d.Minutes())
		Printfln("Seconds: %v", d.Seconds())
		Printfln("Millseconds: %v", d.Milliseconds())
	} else {
		fmt.Println(err.Error())
	}
}
