package main

import (
	"fmt"
	"spamMasker/masking"
)

func main() {
	strTest1 := "http://hehef" //Here's my spammy page: http://hehefouls.netHAHAHA see you. http://sdsd"
	fmt.Println(masking.FindAndMaskLinks(strTest1))
}
