package passgen

import (
	"math/rand"
	"strings"
	"time"
)

var(
	password string
	ex string
	slice = []string{}
	lower = []rune("abcdefghijklmnopqrstuvwxyz")
	upper = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	number = []rune("0123456789")
	symbol = []rune("!@#$%^&*()_+-=}{}[]/?>.,<''`~")
	all = []rune(string(lower)+string(upper)+string(number)+string(symbol))
	lower1 = []rune(string(lower)+string(upper))
	lower2 = []rune(string(lower)+string(number))
	lower3 = []rune(string(lower)+string(symbol))
	lower4 = []rune(string(lower)+string(upper)+string(number))
	lower5 = []rune(string(lower)+string(upper)+string(symbol))
	lower6 = []rune(string(lower)+string(number)+string(symbol))
	upper1 = []rune(string(upper)+string(number))
	upper2 = []rune(string(upper)+string(symbol))
	upper3 = []rune(string(upper)+string(number)+string(symbol))
	number1 = []rune(string(number)+string(symbol))
)

func init(){
	// make unique seed for rand int in every program run
	rand.Seed(time.Now().UTC().UnixNano())
}

func Scramble(r []rune,n int) {

	x := string(r)

	for i := 0; i <= n; i++{
		a := rand.Intn(len(r))
		if strings.Contains(ex,string(x[a])) == true {
			continue
		}
		ex += string(x[a])
	}


}

func Lowercase(n int) string {
	var data = make([]rune, n)

	for i := range data{
		data[i] = lower[rand.Intn(len(lower))] //rand.Intn for random int in 0-maks(n)
	}

	return string(data)
}

func Uppercase(n int) string {
	var data = make([]rune, n)

	for i := range data{
		data[i] = upper[rand.Intn(len(upper))] 
	}

	return string(data)
}

func Numbers(n int) string {
	var data = make([]rune, n)

	for i := range data{
		data[i] = number[rand.Intn(len(number))] 
	}

	return string(data)	
}

func Symbols(n int) string {
	var data = make([]rune, n)

	for i := range data{
		data[i] = symbol[rand.Intn(len(symbol))] 
	}

	return string(data)	
}

func Make(Length int,LowerCase bool,UpperCase bool,Number bool,Symbol bool) string {

	if LowerCase == true && UpperCase == false && Number == false && Symbol == false {
		password = Lowercase(Length)
	}else if LowerCase == false && UpperCase == true && Number == false && Symbol == false{
		password = Uppercase(Length)
	}else if LowerCase == false && UpperCase == false && Number == true && Symbol == false{
		password = Numbers(Length)
	}else if LowerCase == false && UpperCase == false && Number == false && Symbol == true{
		password = Symbols(Length)
	}

	if LowerCase == true && UpperCase == true && Number == false && Symbol == false{
		Scramble(lower1,5)
		password = ex
	}

	return password
}


