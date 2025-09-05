package main

import "fmt"


func makeIntersection(s1, s2 []int) (bool, []int){
	voc := make(map[int]int)

	for _, num := range s1{
		voc[num]++
	}
	for _, num := range s2{
		voc[num]++
	}
	var res []int
	for num, _ := range voc{
		if voc[num] > 1{
			res =  append(res, num)
		}
	}
	if res == nil{
		return false, res
	}
	return true,res
}
func main(){
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	fmt.Println(makeIntersection(a,b))
}