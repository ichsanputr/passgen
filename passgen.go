package passgen

import (
	"math/rand"
	"strings"
	"time"
)

var(
	password string
	myrune []rune
	// all = []rune(string(lower)+string(upper)+string(number)+string(symbol))	
	// lower6 = []rune(string(lower)+string(number)+string(symbol))
	// upper1 = []rune(string(upper)+string(number))
	// upper2 = []rune(string(upper)+string(symbol))
	// upper3 = []rune(string(upper)+string(number)+string(symbol))
	// number1 = []rune(string(number)+string(symbol))
)

func init(){
	// make unique seed for rand int in every program run
	rand.Seed(time.Now().UTC().UnixNano())
}

func Scramble(r []rune,n int) {

	x := string(r)

	for i := 0; i < n; i++{
		a := rand.Intn(len(r))
		if strings.Contains(password,string(x[a])) == true {
			continue
		}
		password += string(x[a])
	}
}

func Lowercase(n int) string {
	var data = make([]rune, n)
	myrune = []rune("abcdefghijklmnopqrstuvwxyz")

	for i := range data{
		data[i] = myrune[rand.Intn(len(myrune))] //rand.Intn for random int in 0-maks(n)
	}

	return string(data)
}

func Uppercase(n int) string {
	var data = make([]rune, n)
	myrune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for i := range data{
		data[i] = myrune[rand.Intn(len(myrune))] 
	}

	return string(data)
}

func Numbers(n int) string {
	var data = make([]rune, n)
	myrune = []rune("0123456789")

	for i := range data{
		data[i] = myrune[rand.Intn(len(myrune))] 
	}

	return string(data)	
}

func Symbols(n int) string {
	var data = make([]rune, n)
	myrune = []rune("!@#$%^&*()_+-=}{}[]/?>.,<''`~")

	for i := range data{
		data[i] = myrune[rand.Intn(len(myrune))] 
	}

	return string(data)	
}

func OptionLower(Length int,LowerCase bool,UpperCase bool,Number bool,Symbol bool){
	if LowerCase == true && UpperCase == true && Number == false && Symbol == false{
		myrune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		Scramble(myrune,Length)
	}

	if LowerCase == true && UpperCase == false && Number == true && Symbol == false {
		myrune = []rune("abcdefghijklmnopqrstuvwxy123456789")
		Scramble(myrune,Length)
	}

	if LowerCase == true && UpperCase == false && Number == false && Symbol == true {
		myrune = []rune("abcdefghijklmnopqrstuvwxy!@#$%^&*()_+-=}{}[]/?>.,<''`~")
		Scramble(myrune,Length)
	}
	
	if LowerCase == true && UpperCase == true && Number == true && Symbol == false {
		myrune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
		Scramble(myrune,Length)
	}

	if LowerCase == true && UpperCase == true && Number == false && Symbol == true {
		myrune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+-=}{}[]/?>.,<''`~")
		Scramble(myrune,Length)
	}

	if LowerCase == true && UpperCase == false && Number == true && Symbol == true {
		myrune = []rune("abcdefghijklmnopqrstuvwxyz123456789!@#$%^&*()_+-=}{}[]/?>.,<''`~")
		Scramble(myrune,Length)
	}
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

	OptionLower(Length,LowerCase,UpperCase,Number,Symbol)

	return password
}


