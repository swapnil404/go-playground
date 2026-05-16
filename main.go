package main

import "fmt"

func main() {
	var cool string = "yuhh"
	if gangcheck(cool) == true {
		fmt.Println("yo gang")
	}else{
		fmt.Println("piss off")
	}
}

func gangcheck(cool string) bool {
	if cool == "yuhh" {
		return true
	}
	return false
}
