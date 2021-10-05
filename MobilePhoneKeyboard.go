package main

import (
	"fmt"
	s "strings"
)

func main() {
	fmt.Println(Solution("aCaBBCCab"))
}

// 1 = CapsLock off, 0 = CapsLock on
func Solution(S string) string {
	arr := s.Split(S, "")
	caplocks := false
	for i := 0; i < len(arr); i++ {
		if arr[i] == "C" {
			arr[i] = "0"
		}
		if arr[i] == "B" {
			arr[i] = "1"
		}
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == "0" {
			caplocks = !caplocks
			arr[i] = "_"
			if caplocks {
				for j := i + 1; j < len(arr); j++ {
					if arr[j] != "0" {
						arr[j] = s.ToUpper(arr[j])
					} else {
						break
					}
				}
			}
		}
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == "1" {
			arr[i] = "_"
			for j := i - 1; j > -1; j-- {
				if arr[j] != "_" {
					arr[j] = "_"
					break
				}
			}
		}
	}
	S = ""
	for i := 0; i < len(arr); i++ {
		if arr[i] != "_" {
			S += arr[i]
		}
	}
	return S
}
