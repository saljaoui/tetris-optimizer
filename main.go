package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("sample.txt")

	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(string(file), "\r\n")
	fmt.Println(data)
	CheckTetris(data)
}
func CheckTetris(data []string) {

	var all []int = []int{ 2}
	// var test int

	for _, trs := range data {
		var cont int
		var ok int = -1
		for i, s := range trs {
			for _, a := range all {
				// if a == 1 {
				// 	if s == '#' {
				// 		if ok == -1 {
				// 			ok = i
				// 		}
				// 		if i == ok {
				// 			cont++
				// 		}
				// 		if i != ok || cont > 4 {
				// 			test ++
				// 			println(test)
				// 		}
				// 	}
				// } else 
				if a == 2 {
					if s == '#' {
						if ok == -1 {
							ok = i
						}
						if i == ok {
							cont++
						}
						if cont > 2 {
							fmt.Println("notOK2")
						}
					}

				}

			}
		}
	}
	fmt.Println("OK")
}
