package main

import (
	"fmt"
	"flag"
	"encoding/hex"
)

func reverseString(data string) string{
	r := []rune(data)
    var res []rune
    for i := len(r) - 1; i >= 0; i-- {
        res = append(res, r[i])
    }
	return string(res)
}

func gen_path(len int) string {
	var pattern string
	p1 := "ABCDEFGHIJKLMNOPQRSTUVWZYZ"
	p2 := "abcdefghijklmnopqrstuvwxyz"
	p3 := "0123456789"
	count := 0
	for count < len{
		for _, c1 := range p1 {
			for _, c2 := range p2 {
				for _, c3 := range p3{
					pattern += fmt.Sprintf("%c%c%c", c1,c2,c3)
					count+=3
					if count >= len {
						return pattern[:len-1]
					}
				}
			}
		}
	}
	return pattern[:len-1]

}

func main(){
	var pattern_len int
	var hex2detect string
	

	flag.IntVar(&pattern_len, "lenght", 1024, "Lenght of the pattern to generate.")
	flag.StringVar(&hex2detect, "find", "", "Data to locate in the pattern")
	flag.Parse()
	
	pattern := gen_path(pattern_len)

	
	if hex2detect == "" {
		fmt.Println(pattern)
	}else{
		byte2detect, convErr := hex.DecodeString(hex2detect)
		if convErr != nil {
			fmt.Println(convErr.Error())
			return
		}
		text2detect := reverseString(string(byte2detect))
		found := false
		text2detectlen := len(text2detect)
		for i:=0; i<pattern_len-text2detectlen; i++{
			if pattern[i:i+(text2detectlen)] == text2detect{
				fmt.Printf("Matching at %d bytes\n", i)
				found = true
			}
		}
		if !found{
			fmt.Println("No match")
		}
	}

}