package passgen

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

var (
	password string
	ex       string
	myrune   []rune
	lower    = []rune("abcdefghijklmnopqrstuvwxyz")
	upper    = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	number   = []rune("0123456789")
	symbol   = []rune("!@#$%^&*()_+-=}{}[]/?>.,<''`~")
)

func init() {
	// make unique seed for rand int in every program run
	rand.Seed(time.Now().UTC().UnixNano())
}

func Scramble(n int, r ...[]rune) {
	s := float64(n / len(r))
	// var x int

	// switch n{
	// case 1:

	// }

	// membagi rune ke ex secara rata
	for _, v := range r {
		// n == genap
		if n%2 == 0 {
			if n % len(r) == 1{
			   // for 8,3 & 10,4 & 6,4
				for i := 0; i < n/len(r); i++ {
					ex += string(v[rand.Intn(len(v))])
				}
			}else if n % len(r) == 0{
				// for 2,2 & 4,4 & 4,2 & 6,3 & 12,3
				for i := 0; i < n/len(r); i++ {
					ex += string(v[rand.Intn(len(v))])
				}
			}else if n < len(r){

			}
		// n == ganjil 
		} else {
			for i := 1; i <= int(math.Round(s)); i++ {
				ex += string(v[rand.Intn(len(v))])
			}
		}
	}

	// rand ex & fill password
	for i := 0; i < len(ex); i++ {
		a := rand.Intn(len(ex))
		if strings.Contains(password, string(ex[a])) == true {
			i--
			continue
		}
		password += string(ex[a])
	}
}

func Lowercase(n int) string {
	var data = make([]rune, n)

	for i := range data {
		data[i] = lower[rand.Intn(len(lower))] //rand.Intn for random int in 0-maks(n)
	}

	return string(data)
}

func Uppercase(n int) string {
	var data = make([]rune, n)

	for i := range data {
		data[i] = upper[rand.Intn(len(upper))]
	}

	return string(data)
}

func Numbers(n int) string {
	var data = make([]rune, n)

	for i := range data {
		data[i] = number[rand.Intn(len(number))]
	}

	return string(data)
}

func Symbols(n int) string {
	var data = make([]rune, n)

	for i := range data {
		data[i] = symbol[rand.Intn(len(symbol))]
	}

	return string(data)
}

func OptionLower(Length int, LowerCase bool, UpperCase bool, Number bool, Symbol bool) {
	if LowerCase == true && UpperCase == true && Number == false && Symbol == false {
		Scramble(Length, lower, upper)
	}

	if LowerCase == true && UpperCase == false && Number == true && Symbol == false {
		Scramble(Length, lower, number)
	}

	if LowerCase == true && UpperCase == false && Number == false && Symbol == true {
		Scramble(Length, lower, symbol)
	}

	if LowerCase == true && UpperCase == true && Number == true && Symbol == false {
		Scramble(Length, lower, upper, number)
	}

	if LowerCase == true && UpperCase == true && Number == false && Symbol == true {
		Scramble(Length, lower, upper, symbol)
	}

	if LowerCase == true && UpperCase == false && Number == true && Symbol == true {
		Scramble(Length, lower, number, symbol)
	}
	
	if LowerCase == true && UpperCase == true && Number == true && Symbol == true {
		Scramble(Length, lower, number, upper,symbol)
	}
}

func Make(Length int, LowerCase bool, UpperCase bool, Number bool, Symbol bool) string {

	if LowerCase == true && UpperCase == false && Number == false && Symbol == false {
		password = Lowercase(Length)
	} else if LowerCase == false && UpperCase == true && Number == false && Symbol == false {
		password = Uppercase(Length)
	} else if LowerCase == false && UpperCase == false && Number == true && Symbol == false {
		password = Numbers(Length)
	} else if LowerCase == false && UpperCase == false && Number == false && Symbol == true {
		password = Symbols(Length)
	}

	OptionLower(Length, LowerCase, UpperCase, Number, Symbol)

	return password
}
