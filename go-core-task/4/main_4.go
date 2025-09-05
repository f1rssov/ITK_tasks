package main

import "fmt"

func makeDifference(s1, s2 []string) []string{
	voc := make(map[string]int)

	for _, str := range s2{
		voc[str]++
	}
	var res []string
	for _, str := range s1{
		if voc[str] == 0{
			res =  append(res, str)
		}
	}
	return res
}

func main(){
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	res := makeDifference(slice1, slice2)
	fmt.Println(res)
}