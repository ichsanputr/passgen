package passgen

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	password string
	ex       string
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

	// rand array rune r
	rand.Shuffle(len(r), func(i, j int) {
		r[i], r[j] = r[j], r[i]
	})

	// declaration variable
	s := (n - (n % len(r))) / len(r)
	residu := n % len(r)

	// divide rune to ex evenly
	for _, v := range r {
		if n%len(r) > 0 {
			// for 8,3 & 11,4 & 6,4 & 9,4
			for i := 0; i < s; i++ {
				ex += string(v[rand.Intn(len(v))])

				if residu != 0 && i == 0 {
					ex += string(v[rand.Intn(len(v))])
					residu--
				}
			}
		} else if n%len(r) == 0 {
			// for  & 9,3 & 4,2 & 15,3 & 12,3
			for i := 0; i < n/len(r); i++ {
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

func Option(Length int, LowerCase bool, UpperCase bool, Number bool, Symbol bool) {
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
		Scramble(Length, lower, number, upper, symbol)
	}
}

func Make(Length int, LowerCase bool, UpperCase bool, Number bool, Symbol bool) string {

	// validate input
	if err := ValidatePass(Length, LowerCase, UpperCase, Number, Symbol); err != nil {
		fmt.Println(err.Error()) // Error() for display error string in errors.New()
	}

	// check condition
	if LowerCase == true && UpperCase == false && Number == false && Symbol == false {
		password = Lowercase(Length)
	} else if LowerCase == false && UpperCase == true && Number == false && Symbol == false {
		password = Uppercase(Length)
	} else if LowerCase == false && UpperCase == false && Number == true && Symbol == false {
		password = Numbers(Length)
	} else if LowerCase == false && UpperCase == false && Number == false && Symbol == true {
		password = Symbols(Length)
	}

	Option(Length, LowerCase, UpperCase, Number, Symbol)

	return password
}
