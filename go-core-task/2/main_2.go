package main
import (
	"math/rand"
	"fmt"
	"time"
)

func sliceExample(src []int) []int{
	var newSlice []int
	for _, n := range src{
		if n % 2 == 0{
			newSlice = append(newSlice, n)
		}
	}
	return newSlice
}

func addElements(src []int, num int) []int {
	return append(src, num)
}

func copySlice(src []int) []int{
	newSlice := make([]int, len(src))
	copy(newSlice, src)
	return newSlice
}

func  removeElement(src []int, ind int) []int{
	if len(src) > 1{
		return append(src[:ind-1], src[ind:]...)
	}
	return src
}

func main(){
	rand.Seed(time.Now().UnixNano())

	var originalSlice []int
	for i:=0; i<10; i++{
		originalSlice =  append(originalSlice, rand.Int())
	}
	////1
	fmt.Println(originalSlice)
	fmt.Println()
	/////2
	fmt.Println(sliceExample(originalSlice))
	fmt.Println()
	/////3
	fmt.Println(addElements(originalSlice, 10))
	fmt.Println()
	/////4
	copied := copySlice(originalSlice)
	originalSlice = append(originalSlice, 10)
	fmt.Println(originalSlice)
	fmt.Println(copied)
	fmt.Println()
	/////5
	fmt.Println(removeElement(originalSlice, 2))
}