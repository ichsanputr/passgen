package passgen

import (
	"fmt"
	"testing"
)

// func TestLowerCase(t *testing.T) {
// 	fmt.Println(Lowercase(7))
// }
// func TestUpperCase(t *testing.T) {
// 	fmt.Println(Uppercase(8))
// 	fmt.Println(rand.Intn(5))
// }
// func TestNumber(t *testing.T) {
// 	fmt.Println(Numbers(8))
// }
// func TestSymbol(t *testing.T) {
// 	fmt.Println(Symbols(8))
// }

func TestMake(t *testing.T){
	fmt.Println(Make(4,true,true,false,false))
}