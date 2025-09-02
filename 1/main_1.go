package main
import (
	"reflect"
	"fmt"
	"strings"
	"encoding/hex"
	"crypto/sha256"
)

func getTypePrint(v interface{}){
	fmt.Println(v, reflect.TypeOf(v).String())
}

func myToString(vars ...interface{}) string{
	parts := make([]string, len(vars))
	for i, v := range vars{
		parts[i] = fmt.Sprintf("%v", v)
	}
	return strings.Join(parts, ", ")
}

func hashRunes(r []rune, salt string) string{
	mid := len(r) / 2
	saltRune:= []rune(salt)

	withSalt := append(r[:mid], append(saltRune, r[mid:]...)...)
	data := []byte(string(withSalt))
	sum:=sha256.Sum256(data)

	return hex.EncodeToString(sum[:])
}

func main(){
////// 1
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A     
	var pi float64 = 3.14             
	var name string = "Golang"         
	var isActive bool = true          
	var complexNum complex64 = 1 + 2i
////// 2
	getTypePrint(numDecimal)
	getTypePrint(numOctal)
	getTypePrint(numHexadecimal)
	getTypePrint(pi)
	getTypePrint(name)
	getTypePrint(isActive)
	getTypePrint(complexNum)
	fmt.Println()
////// 3
	resString := myToString(numDecimal, numOctal, numHexadecimal,
		pi, name, isActive, complexNum)
////// 4
	resRunes := []rune(resString)
	fmt.Println(resRunes)
	for _, c := range resRunes{
		fmt.Printf("%c", c)
	}
	fmt.Println()
////// 5
	resHash := hashRunes(resRunes, "go-2024")
	fmt.Println(resHash)
}


